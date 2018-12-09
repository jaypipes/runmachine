package metadata

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUnknown = status.Errorf(
		codes.Unknown,
		"an unknown error occurred.",
	)
	ErrNotFound = status.Errorf(
		codes.NotFound,
		"object could not be found.",
	)
	ErrSessionUserRequired = status.Errorf(
		codes.FailedPrecondition,
		"user is required in session.",
	)
	ErrSessionPartitionRequired = status.Errorf(
		codes.FailedPrecondition,
		"partition is required in session.",
	)
	ErrSessionProjectRequired = status.Errorf(
		codes.FailedPrecondition,
		"project is required in session.",
	)
	ErrFailedExpandObjectFilters = status.Errorf(
		codes.FailedPrecondition,
		"failed to expand object filters.",
	)
	ErrObjectFilterRequired = status.Errorf(
		codes.FailedPrecondition,
		"object filter is required when fetching object.",
	)
	ErrFailedExpandPropertySchemaFilters = status.Errorf(
		codes.FailedPrecondition,
		"failed to expand property schema filters.",
	)
	ErrMultipleRecordsFound = status.Errorf(
		codes.FailedPrecondition,
		"multiple records found (expected single record match).",
	)
	ErrPartitionUnknown = status.Errorf(
		codes.FailedPrecondition,
		"unknown partition.",
	)
	ErrPartitionRequired = status.Errorf(
		codes.FailedPrecondition,
		"partition is required.",
	)
	ErrObjectTypeRequired = status.Errorf(
		codes.FailedPrecondition,
		"object type is required.",
	)
	ErrPropertyKeyRequired = status.Errorf(
		codes.FailedPrecondition,
		"property key is required.",
	)
	ErrSchemaRequired = status.Errorf(
		codes.FailedPrecondition,
		"schema is required.",
	)
	ErrPropertySchemaFilterRequired = status.Errorf(
		codes.FailedPrecondition,
		"property schema filter is required.",
	)
	ErrPropertySchemaObjectRequired = status.Errorf(
		codes.FailedPrecondition,
		"property schema object is required.",
	)
	ErrSearchRequired = status.Errorf(
		codes.FailedPrecondition,
		"Either UUID or name to search for is required.",
	)
	ErrCodeRequired = status.Errorf(
		codes.FailedPrecondition,
		"A code to search for is required.",
	)
	ErrBootstrapTokenRequired = status.Errorf(
		codes.FailedPrecondition,
		"bootstrap token is required.",
	)
	ErrPartitionNameRequired = status.Errorf(
		codes.FailedPrecondition,
		"partition name is required.",
	)
	ErrAtLeastOneObjectFilterRequired = status.Errorf(
		codes.FailedPrecondition,
		"at least one object filter is required.",
	)
	ErrObjectDeleteFailed = status.Errorf(
		codes.FailedPrecondition,
		"failed to delete object (check response errors collection).",
	)
)

func errPartitionNotFound(partition string) error {
	return status.Errorf(
		codes.FailedPrecondition,
		"Partition %s not found", partition,
	)
}

func errObjectTypeNotFound(objectType string) error {
	return status.Errorf(
		codes.FailedPrecondition,
		"Object type %s not found", objectType,
	)
}