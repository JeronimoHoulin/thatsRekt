module.exports = class Purge1777663632963 {
    name = 'Purge1777663632963'

    async up(db) {
        // Add `purged` flag + audit timestamps for the new PostPurged event.
        // Default false so existing rows back-fill cleanly without a separate
        // backfill step.
        await db.query(`ALTER TABLE "post" ADD "purged" boolean NOT NULL DEFAULT false`)
        await db.query(`ALTER TABLE "post" ADD "purged_at_block" integer`)
        await db.query(`ALTER TABLE "post" ADD "purged_at_timestamp" TIMESTAMP WITH TIME ZONE`)
    }

    async down(db) {
        await db.query(`ALTER TABLE "post" DROP COLUMN "purged_at_timestamp"`)
        await db.query(`ALTER TABLE "post" DROP COLUMN "purged_at_block"`)
        await db.query(`ALTER TABLE "post" DROP COLUMN "purged"`)
    }
}
