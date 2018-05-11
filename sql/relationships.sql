/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : PostgreSQL
 Source Server Version : 100003
 Source Host           : localhost:5432
 Source Catalog        : putong
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 100003
 File Encoding         : 65001

 Date: 11/05/2018 19:50:40
*/


-- ----------------------------
-- Table structure for relationships
-- ----------------------------
DROP TABLE IF EXISTS "public"."relationships";
CREATE TABLE "public"."relationships" (
  "from_user_id" int8 NOT NULL,
  "to_user_id" int8 NOT NULL,
  "add_time" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "state" "public"."relationship" NOT NULL,
  "id" int8 NOT NULL DEFAULT nextval('relationships_id_seq'::regclass)
)
;
ALTER TABLE "public"."relationships" OWNER TO "postgres";

-- ----------------------------
-- Indexes structure for table relationships
-- ----------------------------
CREATE UNIQUE INDEX "idx_user_id" ON "public"."relationships" USING btree (
  "from_user_id" "pg_catalog"."int8_ops" ASC NULLS LAST,
  "to_user_id" "pg_catalog"."int8_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table relationships
-- ----------------------------
ALTER TABLE "public"."relationships" ADD CONSTRAINT "relationships_pkey" PRIMARY KEY ("id");
