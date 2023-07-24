package db

import (
	"context"
)

func Run(context context.Context, pool Pool) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS mailbox (
			id SERIAL PRIMARY KEY,
			senderId BIGINT NOT NULL,
			senderName VARCHAR(255) NOT NULL,
			item JSONB NOT NULL,
			created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			targetId BIGINT NOT NULL,
			targetName VARCHAR(255) NOT NULL,
			message VARCHAR(255) NOT NULL DEFAULT 'You have a new item in your mailbox!'
		)`,
		// `CREATE TABLE IF NOT EXISTS players (
		// 	id SERIAL PRIMARY KEY,
		// 	robloxId BIGINT NOT NULL UNIQUE,
		// 	robloxName VARCHAR(255) NOT NULL,
		// 	secrets BIGINT NOT NULL DEFAULT 0,
		// 	eggs BIGINT NOT NULL DEFAULT 0,
		// 	bubbles BIGINT NOT NULL DEFAULT 0,
		// 	power BIGINT NOT NULL DEFAULT 0,
		// 	robux BIGINT NOT NULL DEFAULT 0,
		// 	playtime BIGINT NOT NULL DEFAULT 0,
		// 	time_saved TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		// )`,
		// `CREATE TABLE IF NOT EXISTS auctions (
		// 	id SERIAL PRIMARY KEY,
		// 	robloxId BIGINT NOT NULL,
		// 	robloxName VARCHAR(255) NOT NULL,
		// 	itemType VARCHAR(255) NOT NULL,
		// 	itemData JSONB NOT NULL,
		// 	startPrice BIGINT NOT NULL,
		// 	priceType VARCHAR(255) NOT NULL,
		// 	listed TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		// 	status VARCHAR(255) NOT NULL DEFAULT 'OPEN'
		// )`,
		// `CREATE TABLE IF NOT EXISTS pets_exist (
		// 	id SERIAL PRIMARY KEY,
		// 	robloxId BIGINT NOT NULL,
		// 	petId VARCHAR(255) NOT NULL,
		// 	petCount BIGINT NOT NULL DEFAULT 0,
		// 	CONSTRAINT uc_robloxid_petid UNIQUE (robloxId, petId)
		// )`,
	}

	for _, query := range queries {
		_, err := pool.Exec(context, query)
		if err != nil {
			return err
		}
	}

	return nil
}
