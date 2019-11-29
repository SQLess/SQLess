/*
 * Copyright 2018-2019 The CovenantSQL Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package internal

import (
	"flag"

	"github.com/SQLess/SQLess/client"
)

// CmdDrop is cql drop command entity.
var CmdDrop = &Command{
	UsageLine: "cql drop [common params] [-wait-tx-confirm] dsn",
	Short:     "drop a database by dsn or database id",
	Long: `
Drop drops a CQL database by DSN or database ID.
e.g.
    cql drop cqlprotocol://4119ef997dedc585bfbcfae00ab6b87b8486fab323a8e107ea1fd4fc4f7eba5c

Since CQL is built on top of blockchains, you may want to wait for the transaction
confirmation before the drop operation takes effect.
e.g.
    cql drop -wait-tx-confirm cqlprotocol://4119ef997dedc585bfbcfae00ab6b87b8486fab323a8e107ea1fd4fc4f7eba5c
`,
	Flag:       flag.NewFlagSet("Drop params", flag.ExitOnError),
	CommonFlag: flag.NewFlagSet("Common params", flag.ExitOnError),
	DebugFlag:  flag.NewFlagSet("Debug params", flag.ExitOnError),
}

func init() {
	CmdDrop.Run = runDrop

	addCommonFlags(CmdDrop)
	addConfigFlag(CmdDrop)
	addWaitFlag(CmdDrop)
}

func runDrop(cmd *Command, args []string) {
	commonFlagsInit(cmd)

	if len(args) != 1 {
		ConsoleLog.Error("drop command need CQL dsn or database_id string as param")
		SetExitStatus(1)
		printCommandHelp(cmd)
		Exit()
	}

	configInit()

	dsn := args[0]

	// drop database
	if _, err := client.ParseDSN(dsn); err != nil {
		// not a dsn/dbid
		ConsoleLog.WithField("db", dsn).WithError(err).Error("not a valid dsn")
		SetExitStatus(1)
		return
	}

	txHash, err := client.Drop(dsn)
	if err != nil {
		// drop database failed
		ConsoleLog.WithField("db", dsn).WithError(err).Error("drop database failed")
		SetExitStatus(1)
		return
	}

	if waitTxConfirmation {
		err = wait(txHash)
		if err != nil {
			ConsoleLog.WithField("db", dsn).WithError(err).Error("drop database failed")
			SetExitStatus(1)
			return
		}
	}

	// drop database success
	ConsoleLog.Infof("drop database %#v success", dsn)
	return
}
