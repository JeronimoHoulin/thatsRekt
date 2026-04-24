// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

/// @title  ThatsRekt
/// @notice Public register of rekt addresses.
///         Any whitelisted address can add entries or manage the whitelist.
///         Removal is a two-step process: one whitelisted address proposes,
///         a different whitelisted address executes.
contract ThatsRekt {
    /*//////////////////////////////////////////////////////////////
                                CONSTANTS
    //////////////////////////////////////////////////////////////*/

    address public constant JERRYTHEKID = 0x9E8680dbBcA1127add812abE209A10E621b385dF;
    address public constant BAUTI       = 0xda1b9dFA299d655135C1ECdc4f0b4c9aED9a7f45;

    /*//////////////////////////////////////////////////////////////
                                 STORAGE
    //////////////////////////////////////////////////////////////*/

    mapping(address => bool) public isWhitelisted;

    address[] private _rektList;
    mapping(address => bool)    private _isRekt;
    mapping(address => uint256) private _rektIndex; // 1-indexed; 0 = absent

    struct RemovalProposal {
        address   proposer;
        bool      executed;
        address[] targets;
    }

    uint256 public proposalCount;
    mapping(uint256 => RemovalProposal) private _proposals;

    /*//////////////////////////////////////////////////////////////
                                 EVENTS
    //////////////////////////////////////////////////////////////*/

    event WhitelistUpdated(address indexed account, bool status);
    event AddedRekt(address indexed account);
    event RemovedRekt(address indexed account);
    event RemovalProposed(uint256 indexed id, address indexed proposer, address[] targets);
    event RemovalExecuted(uint256 indexed id, address indexed executor);

    /*//////////////////////////////////////////////////////////////
                                 ERRORS
    //////////////////////////////////////////////////////////////*/

    error NotWhitelisted();
    error ProposalNotFound();
    error CannotSelfExecute();
    error AlreadyExecuted();

    /*//////////////////////////////////////////////////////////////
                               CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/

    constructor() {
        _setWhitelisted(JERRYTHEKID, true);
        _setWhitelisted(BAUTI, true);
    }

    /*//////////////////////////////////////////////////////////////
                            WHITELIST MANAGEMENT
                         (any whitelisted address)
    //////////////////////////////////////////////////////////////*/

    function addWhitelisted(address account) external onlyWhitelisted {
        _setWhitelisted(account, true);
    }

    function removeWhitelisted(address account) external onlyWhitelisted {
        _setWhitelisted(account, false);
    }

    /*//////////////////////////////////////////////////////////////
                              READ FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns every address currently in the rekt list.
    function isRekt() external view returns (address[] memory) {
        return _rektList;
    }

    /// @notice Returns the full details of a removal proposal.
    function getProposal(uint256 id)
        external
        view
        returns (address proposer, bool executed, address[] memory targets)
    {
        RemovalProposal storage p = _proposals[id];
        return (p.proposer, p.executed, p.targets);
    }

    /*//////////////////////////////////////////////////////////////
                             WRITE FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Add addresses to the rekt list. Caller must be whitelisted.
    ///         Duplicates (within the call or already present) are silently skipped.
    function addRekt(address[] calldata targets) external onlyWhitelisted {
        uint256 len = targets.length;
        for (uint256 i; i < len; ++i) {
            _addToRekt(targets[i]);
        }
    }

    /// @notice Propose the removal of one or more addresses from the rekt list.
    ///         A different whitelisted address must call executeRemoval() to carry it out.
    /// @return id  The proposal ID to pass to executeRemoval().
    function proposeRemoval(address[] calldata targets) external onlyWhitelisted returns (uint256 id) {
        id = proposalCount;
        unchecked { ++proposalCount; }

        RemovalProposal storage p = _proposals[id];
        p.proposer = msg.sender;

        uint256 len = targets.length;
        for (uint256 i; i < len; ++i) {
            p.targets.push(targets[i]);
        }

        emit RemovalProposed(id, msg.sender, targets);
    }

    /// @notice Execute a pending removal proposal. Caller must be whitelisted and
    ///         must not be the same address that proposed it.
    ///         Addresses no longer in the rekt list are silently skipped.
    function executeRemoval(uint256 id) external onlyWhitelisted {
        RemovalProposal storage p = _proposals[id];
        if (p.proposer == address(0)) revert ProposalNotFound();
        if (p.proposer == msg.sender) revert CannotSelfExecute();
        if (p.executed)               revert AlreadyExecuted();

        p.executed = true;

        uint256 len = p.targets.length;
        for (uint256 i; i < len; ++i) {
            _removeFromRekt(p.targets[i]);
        }

        emit RemovalExecuted(id, msg.sender);
    }

    /*//////////////////////////////////////////////////////////////
                               MODIFIERS
    //////////////////////////////////////////////////////////////*/

    modifier onlyWhitelisted() {
        if (!isWhitelisted[msg.sender]) revert NotWhitelisted();
        _;
    }

    /*//////////////////////////////////////////////////////////////
                            INTERNAL HELPERS
    //////////////////////////////////////////////////////////////*/

    function _setWhitelisted(address account, bool status) internal {
        isWhitelisted[account] = status;
        emit WhitelistUpdated(account, status);
    }

    function _addToRekt(address target) internal {
        if (_isRekt[target]) return;
        _isRekt[target] = true;
        _rektList.push(target);
        _rektIndex[target] = _rektList.length; // 1-indexed
        emit AddedRekt(target);
    }

    /// @dev O(1) removal via swap-and-pop using the 1-indexed _rektIndex mapping.
    function _removeFromRekt(address target) internal {
        uint256 idx = _rektIndex[target]; // 1-indexed; 0 = not present
        if (idx == 0) return;

        uint256 lastIdx = _rektList.length;
        if (idx != lastIdx) {
            address last = _rektList[lastIdx - 1];
            _rektList[idx - 1] = last;
            _rektIndex[last]   = idx;
        }
        _rektList.pop();
        delete _rektIndex[target];
        delete _isRekt[target];
        emit RemovedRekt(target);
    }
}
