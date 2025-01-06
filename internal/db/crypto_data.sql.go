// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: crypto_data.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createBTCCryptoData = `-- name: CreateBTCCryptoData :one
INSERT INTO btc_crypto_data(master_pub_key) VALUES ($1)
RETURNING id, master_pub_key, last_major_index, last_minor_index
`

// BTC
func (q *Queries) CreateBTCCryptoData(ctx context.Context, masterPubKey string) (BtcCryptoDatum, error) {
	row := q.db.QueryRow(ctx, createBTCCryptoData, masterPubKey)
	var i BtcCryptoDatum
	err := row.Scan(
		&i.ID,
		&i.MasterPubKey,
		&i.LastMajorIndex,
		&i.LastMinorIndex,
	)
	return i, err
}

const createCryptoData = `-- name: CreateCryptoData :one
INSERT INTO crypto_data(xmr_id, user_id) VALUES ($1, $2)
RETURNING user_id, xmr_id, btc_id
`

type CreateCryptoDataParams struct {
	XmrID  pgtype.UUID
	UserID pgtype.UUID
}

func (q *Queries) CreateCryptoData(ctx context.Context, arg CreateCryptoDataParams) (CryptoDatum, error) {
	row := q.db.QueryRow(ctx, createCryptoData, arg.XmrID, arg.UserID)
	var i CryptoDatum
	err := row.Scan(&i.UserID, &i.XmrID, &i.BtcID)
	return i, err
}

const createXMRCryptoData = `-- name: CreateXMRCryptoData :one
INSERT INTO xmr_crypto_data(priv_view_key, pub_spend_key) VALUES ($1, $2)
RETURNING id, priv_view_key, pub_spend_key, last_major_index, last_minor_index
`

type CreateXMRCryptoDataParams struct {
	PrivViewKey string
	PubSpendKey string
}

// XMR
func (q *Queries) CreateXMRCryptoData(ctx context.Context, arg CreateXMRCryptoDataParams) (XmrCryptoDatum, error) {
	row := q.db.QueryRow(ctx, createXMRCryptoData, arg.PrivViewKey, arg.PubSpendKey)
	var i XmrCryptoDatum
	err := row.Scan(
		&i.ID,
		&i.PrivViewKey,
		&i.PubSpendKey,
		&i.LastMajorIndex,
		&i.LastMinorIndex,
	)
	return i, err
}

const findCryptoDataByUserId = `-- name: FindCryptoDataByUserId :one
SELECT user_id, xmr_id, btc_id FROM crypto_data 
WHERE user_id = $1
`

func (q *Queries) FindCryptoDataByUserId(ctx context.Context, userID pgtype.UUID) (CryptoDatum, error) {
	row := q.db.QueryRow(ctx, findCryptoDataByUserId, userID)
	var i CryptoDatum
	err := row.Scan(&i.UserID, &i.XmrID, &i.BtcID)
	return i, err
}

const findIndicesAndLockBTCCryptoDataById = `-- name: FindIndicesAndLockBTCCryptoDataById :one
SELECT last_major_index, last_minor_index 
FROM btc_crypto_data
WHERE id = $1
FOR UPDATE
`

type FindIndicesAndLockBTCCryptoDataByIdRow struct {
	LastMajorIndex int32
	LastMinorIndex int32
}

func (q *Queries) FindIndicesAndLockBTCCryptoDataById(ctx context.Context, id pgtype.UUID) (FindIndicesAndLockBTCCryptoDataByIdRow, error) {
	row := q.db.QueryRow(ctx, findIndicesAndLockBTCCryptoDataById, id)
	var i FindIndicesAndLockBTCCryptoDataByIdRow
	err := row.Scan(&i.LastMajorIndex, &i.LastMinorIndex)
	return i, err
}

const findIndicesAndLockXMRCryptoDataById = `-- name: FindIndicesAndLockXMRCryptoDataById :one
SELECT last_major_index, last_minor_index 
FROM xmr_crypto_data
WHERE id = $1
FOR UPDATE
`

type FindIndicesAndLockXMRCryptoDataByIdRow struct {
	LastMajorIndex int32
	LastMinorIndex int32
}

func (q *Queries) FindIndicesAndLockXMRCryptoDataById(ctx context.Context, id pgtype.UUID) (FindIndicesAndLockXMRCryptoDataByIdRow, error) {
	row := q.db.QueryRow(ctx, findIndicesAndLockXMRCryptoDataById, id)
	var i FindIndicesAndLockXMRCryptoDataByIdRow
	err := row.Scan(&i.LastMajorIndex, &i.LastMinorIndex)
	return i, err
}

const findKeysAndLockBTCCryptoDataById = `-- name: FindKeysAndLockBTCCryptoDataById :one
SELECT master_pub_key
FROM btc_crypto_data
WHERE id = $1
FOR SHARE
`

func (q *Queries) FindKeysAndLockBTCCryptoDataById(ctx context.Context, id pgtype.UUID) (string, error) {
	row := q.db.QueryRow(ctx, findKeysAndLockBTCCryptoDataById, id)
	var master_pub_key string
	err := row.Scan(&master_pub_key)
	return master_pub_key, err
}

const findKeysAndLockXMRCryptoDataById = `-- name: FindKeysAndLockXMRCryptoDataById :one
SELECT priv_view_key, pub_spend_key
FROM xmr_crypto_data
WHERE id = $1
FOR SHARE
`

type FindKeysAndLockXMRCryptoDataByIdRow struct {
	PrivViewKey string
	PubSpendKey string
}

func (q *Queries) FindKeysAndLockXMRCryptoDataById(ctx context.Context, id pgtype.UUID) (FindKeysAndLockXMRCryptoDataByIdRow, error) {
	row := q.db.QueryRow(ctx, findKeysAndLockXMRCryptoDataById, id)
	var i FindKeysAndLockXMRCryptoDataByIdRow
	err := row.Scan(&i.PrivViewKey, &i.PubSpendKey)
	return i, err
}

const setBTCCryptoDataByUserId = `-- name: SetBTCCryptoDataByUserId :one
UPDATE crypto_data
SET btc_id = $2
WHERE user_id = $1
RETURNING user_id, xmr_id, btc_id
`

type SetBTCCryptoDataByUserIdParams struct {
	UserID pgtype.UUID
	BtcID  pgtype.UUID
}

func (q *Queries) SetBTCCryptoDataByUserId(ctx context.Context, arg SetBTCCryptoDataByUserIdParams) (CryptoDatum, error) {
	row := q.db.QueryRow(ctx, setBTCCryptoDataByUserId, arg.UserID, arg.BtcID)
	var i CryptoDatum
	err := row.Scan(&i.UserID, &i.XmrID, &i.BtcID)
	return i, err
}

const setXMRCryptoDataByUserId = `-- name: SetXMRCryptoDataByUserId :one
UPDATE crypto_data
SET xmr_id = $2 
WHERE user_id = $1
RETURNING user_id, xmr_id, btc_id
`

type SetXMRCryptoDataByUserIdParams struct {
	UserID pgtype.UUID
	XmrID  pgtype.UUID
}

func (q *Queries) SetXMRCryptoDataByUserId(ctx context.Context, arg SetXMRCryptoDataByUserIdParams) (CryptoDatum, error) {
	row := q.db.QueryRow(ctx, setXMRCryptoDataByUserId, arg.UserID, arg.XmrID)
	var i CryptoDatum
	err := row.Scan(&i.UserID, &i.XmrID, &i.BtcID)
	return i, err
}

const updateIndicesBTCCryptoDataById = `-- name: UpdateIndicesBTCCryptoDataById :one
UPDATE btc_crypto_data
SET last_major_index = $2,
    last_minor_index = $3
WHERE id = $1
RETURNING id, master_pub_key, last_major_index, last_minor_index
`

type UpdateIndicesBTCCryptoDataByIdParams struct {
	ID             pgtype.UUID
	LastMajorIndex int32
	LastMinorIndex int32
}

func (q *Queries) UpdateIndicesBTCCryptoDataById(ctx context.Context, arg UpdateIndicesBTCCryptoDataByIdParams) (BtcCryptoDatum, error) {
	row := q.db.QueryRow(ctx, updateIndicesBTCCryptoDataById, arg.ID, arg.LastMajorIndex, arg.LastMinorIndex)
	var i BtcCryptoDatum
	err := row.Scan(
		&i.ID,
		&i.MasterPubKey,
		&i.LastMajorIndex,
		&i.LastMinorIndex,
	)
	return i, err
}

const updateIndicesXMRCryptoDataById = `-- name: UpdateIndicesXMRCryptoDataById :one
UPDATE xmr_crypto_data
SET last_major_index = $2,
    last_minor_index = $3
WHERE id = $1
RETURNING id, priv_view_key, pub_spend_key, last_major_index, last_minor_index
`

type UpdateIndicesXMRCryptoDataByIdParams struct {
	ID             pgtype.UUID
	LastMajorIndex int32
	LastMinorIndex int32
}

func (q *Queries) UpdateIndicesXMRCryptoDataById(ctx context.Context, arg UpdateIndicesXMRCryptoDataByIdParams) (XmrCryptoDatum, error) {
	row := q.db.QueryRow(ctx, updateIndicesXMRCryptoDataById, arg.ID, arg.LastMajorIndex, arg.LastMinorIndex)
	var i XmrCryptoDatum
	err := row.Scan(
		&i.ID,
		&i.PrivViewKey,
		&i.PubSpendKey,
		&i.LastMajorIndex,
		&i.LastMinorIndex,
	)
	return i, err
}

const updateKeysBTCCryptoDataById = `-- name: UpdateKeysBTCCryptoDataById :one
UPDATE btc_crypto_data
SET master_pub_key = $2,
    last_major_index = 0,
    last_minor_index = 0
WHERE id = $1
RETURNING id, master_pub_key, last_major_index, last_minor_index
`

type UpdateKeysBTCCryptoDataByIdParams struct {
	ID           pgtype.UUID
	MasterPubKey string
}

func (q *Queries) UpdateKeysBTCCryptoDataById(ctx context.Context, arg UpdateKeysBTCCryptoDataByIdParams) (BtcCryptoDatum, error) {
	row := q.db.QueryRow(ctx, updateKeysBTCCryptoDataById, arg.ID, arg.MasterPubKey)
	var i BtcCryptoDatum
	err := row.Scan(
		&i.ID,
		&i.MasterPubKey,
		&i.LastMajorIndex,
		&i.LastMinorIndex,
	)
	return i, err
}

const updateKeysXMRCryptoDataById = `-- name: UpdateKeysXMRCryptoDataById :one
UPDATE xmr_crypto_data
SET priv_view_key = $2,
    pub_spend_key = $3,
    last_major_index = 0,
    last_minor_index = 0
WHERE id = $1
RETURNING id, priv_view_key, pub_spend_key, last_major_index, last_minor_index
`

type UpdateKeysXMRCryptoDataByIdParams struct {
	ID          pgtype.UUID
	PrivViewKey string
	PubSpendKey string
}

func (q *Queries) UpdateKeysXMRCryptoDataById(ctx context.Context, arg UpdateKeysXMRCryptoDataByIdParams) (XmrCryptoDatum, error) {
	row := q.db.QueryRow(ctx, updateKeysXMRCryptoDataById, arg.ID, arg.PrivViewKey, arg.PubSpendKey)
	var i XmrCryptoDatum
	err := row.Scan(
		&i.ID,
		&i.PrivViewKey,
		&i.PubSpendKey,
		&i.LastMajorIndex,
		&i.LastMinorIndex,
	)
	return i, err
}
