// Copyright 2022 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
// Copyright 2022 ZORALab Enterprise <tech@zoralab.com>
//
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
//                  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

package hestiaERROR

const (
	// Common errors and operations
	OK Error = iota
	BAD_DESCRIPTOR
	BAD_EXCHANGE
	BAD_EXEC
	BAD_MOUNT
	BAD_PIPE
	BAD_REQUEST
	BAD_STREAM_PIPE
	CANCELED
	CLEANING_REQUIRED
	DEADLOCK
	EXPIRED
	ILLEGAL_BYTE_SEQUENCE
	ILLEGAL_SEEK
	INVALID_ARGUMENT
	IS_EMPTY
	MAXED_EXCHANGE
	MAXED_QUOTA
	MISSING_LOCK
	NOT_EMPTY
	NOT_PERMITTED
	NOT_POSSIBLE
	NOT_POSSIBLE_BY_RFKILL
	NOT_RECOVERABLE
	OUT_OF_RANGE
	PERMISSION_DENIED
	TIMEOUT
	TOO_MANY_READ
	TOO_MANY_LOOP
	TOO_MANY_REFERENCES
	TOO_MANY_LINK
	TRY_AGAIN
	UNSUPPORTED
	WOULD_BLOCK

	// lifecycle states
	RESTART
	RESUME
	SHUTDOWN
	SLEEP
	STALLED
	STANDBY

	// progress
	PROGRESS_SCHEDULED
	PROGRESS_ALREADY_EXECUTING
	PROGRESS_EXECUTING
	PROGRESS_COMPLETED

	// tri-tier inter-package communications
	LV1_NOT_SYNC
	LV1_PAUSED
	LV1_RESET

	LV2_NOT_SYNC
	LV2_PAUSED
	LV2_RESET

	LV3_NOT_SYNC
	LV3_PAUSED
	LV3_RESET

	// data (input/output parameters, type, etc)
	DATA_BAD
	DATA_EMPTY
	DATA_INVALID
	DATA_IS_UNIQUE
	DATA_MISSING
	DATA_NOT_UNIQUE
	DATA_OVERFLOW
	DATA_REMOVED
	DATA_TOO_LONG
	DATA_MISMATCHED

	// entity (device, file, directory, etc)
	ENTITY_BAD
	ENTITY_BUSY
	ENTITY_DEAD
	ENTITY_EXISTS
	ENTITY_FAULTY
	ENTITY_MISSING
	ENTITY_MISSING_CHILD
	ENTITY_OUT_OF_BUFFER
	ENTITY_POISONED
	ENTITY_READ_ONLY
	ENTITY_TOO_BIG
	ENTITY_TOO_MANY_OPENED
	ENTITY_UNATTACHED

	ENTITY_IS_DIRECTORY
	ENTITY_IS_FILE
	ENTITY_IS_LINK
	ENTITY_IS_SOCKET

	ENTITY_IS_NOT_DIRECTORY
	ENTITY_IS_NOT_FILE
	ENTITY_IS_NOT_LINK
	ENTITY_IS_NOT_SOCKET

	ENTITY_REMOTE_CHANGED
	ENTITY_REMOTE_ERROR
	ENTITY_REMOTE_IO

	ENTITY_MISSING_STREAMABLE_RESOURCES
	ENTITY_NOT_STREAMABLE
	ENTITY_STREAMABLE

	ENTITY_A_TYPEWRITER
	ENTITY_NOT_A_TYPEWRITER

	ENTITY_BAD_DESCRIPTOR
	ENTITY_FILETABLE_OVERFLOW

	// key (cryptography)
	KEY_BAD
	KEY_DESTROYED
	KEY_EXPIRED
	KEY_MISSING
	KEY_REJECTED
	KEY_REVOKED

	// library
	LIBRARY_BAD
	LIBRARY_CORRUPTED
	LIBRARY_EXEC_FAILED
	LIBRARY_MAXED
	LIBRARY_MISSING

	// network
	NETWORK_BAD
	NETWORK_BAD_AD
	NETWORK_DOWN
	NETWORK_NOT_CONNECTED
	NETWORK_RESET
	NETWORK_RFS
	NETWORK_UNREACHABLE

	NETWORK_HOST_DOWN
	NETWORK_HOST_UNREACHABLE
	NETWORK_SOCKET_UNSUPPORTED

	NETWORK_ADDRESS_IN_USE
	NETWORK_ADDRESS_UNAVAILABLE

	NETWORK_CONN_ABORTED
	NETWORK_CONN_IS_CONNECTED
	NETWORK_CONN_MISSING_DEST_ADDRESS
	NETWORK_CONN_MULTIHOP
	NETWORK_CONN_NOT_CONNECTED
	NETWORK_CONN_REFUSED
	NETWORK_CONN_RESET

	NETWORK_PAYLOAD_BAD
	NETWORK_PAYLOAD_EMPTY
	NETWORK_PAYLOAD_MISSING
	NETWORK_PAYLOAD_TOO_LONG

	// protocol
	PROTOCOL_ADDRESS_UNSUPPORTED
	PROTOCOL_BAD
	PROTOCOL_FAMILY_UNSUPPORTED
	PROTOCOL_MISSING
	PROTOCOL_UNSUPPORTED

	// system (e.g. os, interactable system)
	SYSTEM_BAD_IO
	SYSTEM_DEVICE_CROSS_LINK
	SYSTEM_INTERRUPT_CALL
	SYSTEM_INVALID
	SYSTEM_MISSING_BLOCK_DEVICE
	SYSTEM_MISSING_DEVICE
	SYSTEM_MISSING_IO
	SYSTEM_MISSING_PROCESS
	SYSTEM_OUT_OF_DOMAIN
	SYSTEM_OUT_OF_MEMORY
	SYSTEM_OUT_OF_SPACE
	SYSTEM_READ_ONLY_FILESYSTEM

	// user
	USER_ACCESS_BANNED
	USER_ACCESS_LOCKED
	USER_ACCESS_NOT_VERIFIED
	USER_ACCESS_REJECTED
	USER_ACCESS_REVOKED

	USER_ID_BAD
	USER_ID_EXISTS
	USER_ID_MISSING
	USER_MFA_BAD
	USER_MFA_EXPIRED
	USER_MFA_MISSING
	USER_PASSWORD_BAD
	USER_PASSWORD_EXPIRED
	USER_PASSWORD_MISSING
	USER_KEYFILE_BAD
	USER_KEYFILE_EXPIRED
	USER_KEYFILE_MISSING

	USER_ACCESS_TOKEN_BAD
	USER_ACCESS_TOKEN_EXPIRED
	USER_ACCESS_TOKEN_MISSING
	USER_ACCESS_TOKEN_REVOKED
)
