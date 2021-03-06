package dbbuildc

import (
	"fmt"
	"strings"

	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/db"
)

func Run(a app.App) error {
	var deps struct {
		Database        services.Database `dependency:"DatabaseService"`
		DependencyScope app.Scope         `dependency:"DependencyScope"`
		DSQL            db.DSQL           `dependency:"DSQL"`
	}
	if err := a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	tx, err := deps.Database.TX()
	if err != nil {
		return err
	}
	keys, err := deps.DependencyScope.Keys()
	if err != nil {
		return err
	}
	for _, key := range keys {
		if strings.HasSuffix(key, "Table") {
			tableIns, err := deps.DependencyScope.Get(key)
			if err != nil {
				return err
			}
			table, ok := tableIns.(db.Table)
			if !ok {
				return fmt.Errorf("%s is not instance of db.Table", key)
			}
			query, err := deps.DSQL.NewCreateSQL(table.Name(), table.Types())
			if err != nil {
				return err
			}
			tx.MustExec(query)
			fmt.Printf("\n\n%s:\n%s", key, query)
		}
	}
	fmt.Printf("\n\ncreated all tables\n")
	fmt.Printf("commited... ")
	if err := tx.Commit(); err != nil {
		fmt.Printf("fail\n")
		return err
	}
	fmt.Printf("ok\n")
	return nil
}
