// Type augmentations for `otomato-sdk@2.0.557`.
//
// The published `.d.ts` is stale relative to the runtime — SPLIT,
// HTTP_REQUEST, AI, IF, X_POST_TRIGGER, and SEND_EMAIL are present at
// runtime but absent from the typedefs. We declare only the entries
// this package needs and re-export ACTIONS / TRIGGERS through the
// augmented type so every other module imports from here, not directly
// from `otomato-sdk`.

import {
  ACTIONS as RawACTIONS,
  TRIGGERS as RawTRIGGERS,
  type Parameter,
} from 'otomato-sdk';

export interface BlockDescriptor {
  readonly name: string;
  readonly type: number;
  readonly description: string;
  readonly blockId: number;
  readonly image: string;
  readonly parameters: Parameter[];
  readonly output?: { readonly [key: string]: string };
}

export interface ActionsAugmented {
  readonly CORE: {
    readonly SPLIT: { readonly SPLIT: BlockDescriptor };
    readonly CONDITION: { readonly IF: BlockDescriptor };
    readonly HTTP_REQUEST: { readonly HTTP_REQUEST: BlockDescriptor };
  };
  readonly AI: {
    readonly AI: { readonly AI: BlockDescriptor };
  };
  readonly NOTIFICATIONS: {
    readonly EMAIL: { readonly SEND_EMAIL: BlockDescriptor };
  };
}

export interface TriggersAugmented {
  readonly SOCIALS: {
    readonly X: {
      readonly X_POST_TRIGGER: BlockDescriptor;
    };
  };
}

export const ACTIONS = RawACTIONS as unknown as ActionsAugmented;
export const TRIGGERS = RawTRIGGERS as unknown as TriggersAugmented;
