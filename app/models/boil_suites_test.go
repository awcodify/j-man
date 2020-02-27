// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Migrations", testMigrations)
	t.Run("Rounds", testRounds)
}

func TestDelete(t *testing.T) {
	t.Run("Migrations", testMigrationsDelete)
	t.Run("Rounds", testRoundsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Migrations", testMigrationsQueryDeleteAll)
	t.Run("Rounds", testRoundsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Migrations", testMigrationsSliceDeleteAll)
	t.Run("Rounds", testRoundsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Migrations", testMigrationsExists)
	t.Run("Rounds", testRoundsExists)
}

func TestFind(t *testing.T) {
	t.Run("Migrations", testMigrationsFind)
	t.Run("Rounds", testRoundsFind)
}

func TestBind(t *testing.T) {
	t.Run("Migrations", testMigrationsBind)
	t.Run("Rounds", testRoundsBind)
}

func TestOne(t *testing.T) {
	t.Run("Migrations", testMigrationsOne)
	t.Run("Rounds", testRoundsOne)
}

func TestAll(t *testing.T) {
	t.Run("Migrations", testMigrationsAll)
	t.Run("Rounds", testRoundsAll)
}

func TestCount(t *testing.T) {
	t.Run("Migrations", testMigrationsCount)
	t.Run("Rounds", testRoundsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Migrations", testMigrationsHooks)
	t.Run("Rounds", testRoundsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Migrations", testMigrationsInsert)
	t.Run("Migrations", testMigrationsInsertWhitelist)
	t.Run("Rounds", testRoundsInsert)
	t.Run("Rounds", testRoundsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Migrations", testMigrationsReload)
	t.Run("Rounds", testRoundsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Migrations", testMigrationsReloadAll)
	t.Run("Rounds", testRoundsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Migrations", testMigrationsSelect)
	t.Run("Rounds", testRoundsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Migrations", testMigrationsUpdate)
	t.Run("Rounds", testRoundsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Migrations", testMigrationsSliceUpdateAll)
	t.Run("Rounds", testRoundsSliceUpdateAll)
}