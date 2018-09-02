DROP SCHEMA IF EXISTS test_resources;
CREATE SCHEMA test_resources;
USE test_resources;

CREATE TABLE resource_classes (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY
, code VARCHAR(200) NOT NULL
, description TEXT CHARACTER SET utf8 COLLATE utf8_bin NULL
, UNIQUE INDEX uix_code (code)
) CHARACTER SET latin1 COLLATE latin1_bin;

CREATE TABLE traits (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY
, code VARCHAR(200) NOT NULL
, description TEXT CHARACTER SET utf8 COLLATE utf8_bin NULL
, UNIQUE INDEX uix_code (code)
) CHARACTER SET latin1 COLLATE latin1_bin;

CREATE TABLE distance_types (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY
, code VARCHAR(200) NOT NULL
, description TEXT CHARACTER SET utf8 COLLATE utf8_bin NULL
, generation INT NOT NULL
, UNIQUE INDEX uix_code (code)
) CHARACTER SET latin1 COLLATE latin1_bin;

CREATE TABLE distances (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY
, type_id INT NOT NULL
, code VARCHAR(200) NOT NULL
, description TEXT CHARACTER SET utf8 COLLATE utf8_bin NULL
, position INT NOT NULL
, generation INT NOT NULL
, UNIQUE INDEX uix_type_code (type_id, code)
) CHARACTER SET latin1 COLLATE latin1_bin;

CREATE TABLE providers (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY
, uuid CHAR(32) NOT NULL
, generation INT UNSIGNED NOT NULL
, root_provider_id INT NOT NULL
, parent_provider_id INT NULL
, nested_left INT NOT NULL
, nested_right INT NOT NULL
, UNIQUE INDEX uix_uuid (uuid)
, UNIQUE INDEX uix_nested_sets (root_provider_id, nested_left, nested_right)
, INDEX ix_parent_provider_id (parent_provider_id)
) CHARACTER SET latin1 COLLATE latin1_bin;

CREATE TABLE inventories (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY
, provider_id BIGINT NOT NULL
, resource_class_id INT NOT NULL
, total BIGINT UNSIGNED NOT NULL
, reserved BIGINT UNSIGNED NOT NULL
, min_unit BIGINT UNSIGNED NOT NULL
, max_unit BIGINT UNSIGNED NOT NULL
, step_size BIGINT UNSIGNED NOT NULL
, allocation_ratio FLOAT NOT NULL
, UNIQUE INDEX uix_provider_resource_class (provider_id, resource_class_id)
, INDEX ix_resource_class_total (resource_class_id, total)
) CHARACTER SET latin1 COLLATE latin1_bin;

CREATE TABLE provider_groups (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY
, uuid CHAR(32) NOT NULL
, region_uuid CHAR(32) NOT NULL
, UNIQUE INDEX uix_uuid (uuid)
, INDEX ix_region_uuid (region_uuid)
) CHARACTER SET latin1 COLLATE latin1_bin;

CREATE TABLE provider_group_members (
  provider_group_id INT NOT NULL
, root_provider_id BIGINT NOT NULL
, PRIMARY KEY (provider_group_id, root_provider_id)
, INDEX (root_provider_id)
);

CREATE TABLE provider_group_distances (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY
, left_provider_group_id INT NOT NULL
, right_provider_group_id INT NOT NULL
, distance_id BIGINT NOT NULL
, UNIQUE INDEX uix_left_right_distance (left_provider_group_id, right_provider_group_id, distance_id)
);

CREATE TABLE consumer_types (
  id SMALLINT NOT NULL AUTO_INCREMENT PRIMARY KEY
, code VARCHAR(200) NOT NULL
, description TEXT CHARACTER SET utf8 COLLATE utf8_bin NULL
, UNIQUE INDEX uix_code (code)
) CHARACTER SET latin1 COLLATE latin1_bin;

CREATE TABLE consumers (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY
, type_id SMALLINT NOT NULL
, uuid CHAR(32) NOT NULL
, generation INT UNSIGNED NOT NULL
, owner_account_uuid CHAR(32) NOT NULL
, owner_project_uuid CHAR(32) NOT NULL
, owner_user_uuid CHAR(32) NOT NULL
, UNIQUE INDEX uix_uuid (uuid)
, INDEX ix_type_id (type_id)
, INDEX ix_owner (owner_project_uuid, owner_user_uuid)
, INDEX ix_account_uuid (owner_account_uuid)
) CHARACTER SET latin1 COLLATE latin1_bin;

CREATE TABLE allocations (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY
, consumer_id BIGINT NOT NULL
, start_time DATETIME NOT NULL
, end_time DATETIME NOT NULL
, INDEX ix_consumer (consumer_id)
, INDEX ix_window (start_time, end_time)
) CHARACTER SET latin1 COLLATE latin1_bin;

CREATE TABLE allocation_items (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY
, allocation_id BIGINT NOT NULL
, provider_id BIGINT NOT NULL
, resource_class_id INT NOT NULL
, used BIGINT UNSIGNED NOT NULL
, INDEX ix_allocation_provider_resource_class (allocation_id, provider_id, resource_class_id)
, INDEX ix_resource_class_provider_allocation (resource_class_id, provider_id, allocation_id)
) CHARACTER SET latin1 COLLATE latin1_bin;
