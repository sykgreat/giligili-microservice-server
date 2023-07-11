package Snowflake

import "testing"

var snowflake, _ = NewSnowflake(1, 1, 1)

func TestSnowflake_NextVal(t *testing.T) {
	t.Log(snowflake.NextVal())
}

func TestSnowflake_ValidateSnowflake(t *testing.T) {
	t.Log(snowflake.ValidateSnowflake())
}
