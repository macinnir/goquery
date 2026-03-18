package goquery_test

import (
	"fmt"
	"testing"

	goquery "github.com/macinnir/query"
	"github.com/macinnir/query/testassets"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQuerySelect(t *testing.T) {
	q := goquery.Select(&testassets.Comment{})
	var e error

	sql, e := q.String()
	require.Nil(t, e)
	assert.Equal(t, "SELECT `t`.* FROM `Comment` `t`", sql)

	sql, e = goquery.Select(&testassets.Comment{}).
		Where(
			goquery.GT("DateCreated", 2),
			goquery.Or(),
			goquery.EQ("Content", "foo's"),
			goquery.Or(),
			goquery.EQ("Name", "bar"),
		).String()
	require.Nil(t, e)
	assert.Equal(t, "SELECT `t`.* FROM `Comment` `t` WHERE `t`.`DateCreated` > 2 OR `t`.`Content` = '"+`foo\'`+"s' OR `t`.`Name` = 'bar'", sql)

	sql, e = goquery.Select(&testassets.Comment{}).
		Where(
			goquery.GT("DateCreated", 2),
			goquery.Or(),
			goquery.EQ("Content", "foo"),
			goquery.And(
				goquery.GTOE("DateCreated", 1),
				goquery.Or(),
				goquery.LTOE("DateCreated", 2),
				goquery.Or(),
				goquery.LT("DateCreated", 3),
			),
		).
		String()
	require.Nil(t, e)
	assert.Equal(t, "SELECT `t`.* FROM `Comment` `t` WHERE `t`.`DateCreated` > 2 OR `t`.`Content` = 'foo' AND ( `t`.`DateCreated` >= 1 OR `t`.`DateCreated` <= 2 OR `t`.`DateCreated` < 3 )", sql)

	sql, e = goquery.Select(&testassets.Comment{}).
		Where(
			goquery.WhereAll(),
			goquery.And(
				goquery.GT("DateCreated", 2),
				goquery.Or(),
				goquery.EQ("Content", "foo"),
			),
		).String()
	require.Nil(t, e)
	assert.Equal(t, "SELECT `t`.* FROM `Comment` `t` WHERE 1=1 AND ( `t`.`DateCreated` > 2 OR `t`.`Content` = 'foo' )", sql)

	sql, e = goquery.Select(&testassets.Comment{}).
		Where(
			goquery.WhereAll(),
			goquery.Or(
				goquery.GT("DateCreated", 2),
				goquery.And(),
				goquery.EQ("Content", "foo"),
			),
			goquery.And(
				goquery.Between("ObjectID", 1, 2),
			),
			goquery.And(),
			goquery.IN("Content", "foo", "bar", "baz"),
			goquery.And(),
			goquery.NE("Content", "quux"),
			goquery.And(),
			goquery.NE("ObjectID", "5"),
		).
		OrderBy("Content", goquery.OrderDirASC).
		Limit(1, 2).String()
	require.Nil(t, e)
	assert.Equal(t, "SELECT `t`.* FROM `Comment` `t` WHERE 1=1 OR ( `t`.`DateCreated` > 2 AND `t`.`Content` = 'foo' ) AND ( `t`.`ObjectID` BETWEEN 1 AND 2 ) AND `t`.`Content` IN ( 'foo', 'bar', 'baz' ) AND `t`.`Content` <> 'quux' AND `t`.`ObjectID` <> 5 ORDER BY `t`.`Content` ASC LIMIT 1 OFFSET 2", sql)
}

func TestQuerySelect_LimitPage(t *testing.T) {
	sql, e := goquery.Select(&testassets.Comment{}).LimitPage(10, 5).String()
	require.Nil(t, e)
	assert.Equal(t, "SELECT `t`.* FROM `Comment` `t` LIMIT 10 OFFSET 50", sql, "LimitPage() should have an offset that multiplies the limit by the page")
}

func TestMultipleOrderBy(t *testing.T) {
	q, e := goquery.Select(&testassets.Comment{}).OrderBy("CommentID", goquery.OrderDirASC).OrderBy("DateCreated", goquery.OrderDirDESC).String()
	assert.Nil(t, e)
	assert.Equal(t, "SELECT `t`.* FROM `Comment` `t` ORDER BY `t`.`CommentID` ASC, `t`.`DateCreated` DESC", q)
}

func TestQuerySelect_InvalidOrderByColumn(t *testing.T) {

	q, e := goquery.Select(&testassets.Comment{}).OrderBy("CommentID", goquery.OrderDirASC).String()
	assert.Nil(t, e)
	assert.Equal(t, "SELECT `t`.* FROM `Comment` `t` ORDER BY `t`.`CommentID` ASC", q)

	q, e = goquery.Select(&testassets.Comment{}).OrderBy("foo", goquery.OrderDirASC).String()
	assert.Equal(t, "Invalid Column Name at ORDER BY in model `Comment` -- foo", e.Error())
	assert.Equal(t, "SELECT `t`.* FROM `Comment` `t` ORDER BY `t`.`foo` ASC", q)
}

func TestQuerySelect_WhereIN(t *testing.T) {
	sql, e := goquery.Select(&testassets.Comment{}).
		Where(
			goquery.IN("Content", "foo", "bar", "baz"),
		).String()
	require.Nil(t, e)
	assert.Equal(t, "SELECT `t`.* FROM `Comment` `t` WHERE `t`.`Content` IN ( 'foo', 'bar', 'baz' )", sql)
}

func TestQuerySelect_WhereNotIN(t *testing.T) {
	sql, e := goquery.Select(&testassets.Comment{}).
		Where(
			goquery.NOTIN("Content", "foo", "bar", "baz"),
		).String()
	require.Nil(t, e)
	assert.Equal(t, "SELECT `t`.* FROM `Comment` `t` WHERE `t`.`Content` NOT IN ( 'foo', 'bar', 'baz' )", sql)
}

func TestQuerySelect_InvalidFieldName(t *testing.T) {

	sql, e := goquery.Select(&testassets.Comment{}).
		Where(
			goquery.EQ("Foo", "Bar"),
		).String()

	require.NotNil(t, e)
	assert.Equal(t, "Invalid Column Name at WHERE... in model `Comment` -- Foo", e.Error())
	assert.Equal(t, "SELECT `t`.* FROM `Comment` `t` WHERE `t`.`Foo` = Bar", sql)

}

func TestQuery_INString(t *testing.T) {

	args := []string{"foo", "bar", "baz"}

	sql, e := goquery.Select(&testassets.Comment{}).
		Where(
			goquery.INString(
				"Content",
				args,
			),
		).String()

	require.Nil(t, e)
	assert.Equal(t, "SELECT `t`.* FROM `Comment` `t` WHERE `t`.`Content` IN ( 'foo', 'bar', 'baz' )", sql)

}

func TestQuery_INInt64(t *testing.T) {

	args := []int64{1, 2, 3}

	sql, e := goquery.Select(&testassets.Comment{}).
		Where(
			goquery.INInt64(
				"CommentID",
				args,
			),
		).String()

	require.Nil(t, e)
	assert.Equal(t, "SELECT `t`.* FROM `Comment` `t` WHERE `t`.`CommentID` IN ( 1, 2, 3 )", sql)

}

func TestQuery_INInt(t *testing.T) {

	args := []int{1, 2, 3}

	sql, e := goquery.Select(&testassets.Comment{}).
		Where(
			goquery.INInt(
				"CommentID",
				args,
			),
		).String()

	require.Nil(t, e)
	assert.Equal(t, "SELECT `t`.* FROM `Comment` `t` WHERE `t`.`CommentID` IN ( 1, 2, 3 )", sql)

}

func TestQuerySelect_EmptyWhereClause(t *testing.T) {

	q := goquery.Select(&testassets.Comment{})
	// TODO extra where clause
	wheres := []*goquery.WherePart{}
	sql, e := q.Where(wheres...).String()
	require.Nil(t, e)
	// assert.Equal(t, "Empty where clause at WHERE in model `Comment` -- ", e.Error())
	assert.Equal(t, "SELECT `t`.* FROM `Comment` `t`", sql)

	q = goquery.Select(&testassets.Comment{})
	sql, e = q.Where(goquery.EQ("CommentID", 1)).String()
	require.Nil(t, e)
	assert.Equal(t, "SELECT `t`.* FROM `Comment` `t` WHERE `t`.`CommentID` = 1", sql)
}

func TestQuerySelect_InvalidField(t *testing.T) {
	sql, e := goquery.Select(&testassets.Comment{}).
		Field("Foo").
		String()

	assert.Equal(t, "SELECT `t`.`Foo` FROM `Comment` `t`", sql)
	require.NotNil(t, e)
	assert.Equal(t, "Invalid Column Name at SELECT...Field in model `Comment` -- Foo", e.Error())
}

func TestWhereLike(t *testing.T) {
	sql, e := goquery.Select(&testassets.Comment{}).Where(goquery.Like("Name", "Foo%")).String()
	require.Nil(t, e)
	assert.Equal(t, "SELECT `t`.* FROM `Comment` `t` WHERE `t`.`Name` LIKE 'Foo%'", sql)
}

func TestWhereLike_InvalidValue(t *testing.T) {
	_, e := goquery.Select(&testassets.Comment{}).Where(goquery.Like("CommentID", "Foo%")).String()
	require.NotNil(t, e)
	assert.Equal(t, "Invalid value at WHERE...LIKE in model `Comment` -- `%d` value: Foo%", e.Error())
}

func TestWhereNotLike(t *testing.T) {
	sql, e := goquery.Select(&testassets.Comment{}).Where(goquery.NotLike("Name", "Foo%")).String()
	require.Nil(t, e)
	assert.Equal(t, "SELECT `t`.* FROM `Comment` `t` WHERE `t`.`Name` NOT LIKE 'Foo%'", sql)
}

func TestWhereNotLike_InvalidValue(t *testing.T) {
	_, e := goquery.Select(&testassets.Comment{}).Where(goquery.NotLike("CommentID", "Foo%")).String()
	require.NotNil(t, e)
	assert.Equal(t, "Invalid value at WHERE...NOT LIKE in model `Comment` -- `%d` value: Foo%", e.Error())
}

func TestUnion(t *testing.T) {
	var e error
	sql, e := goquery.Union(
		goquery.Select(&testassets.Comment{}).Where(goquery.EQ("Content", "bar")),
		goquery.Select(&testassets.Comment{}).Where(goquery.EQ("Content", "baz")),
	)
	require.Nil(t, e)
	assert.Equal(t, "SELECT `t`.* FROM `Comment` `t` WHERE `t`.`Content` = 'bar' UNION ALL SELECT `t`.* FROM `Comment` `t` WHERE `t`.`Content` = 'baz'", sql)
}

func TestUpdate(t *testing.T) {
	sql, e := goquery.Update(&testassets.Comment{}).
		Set("Content", "bar").
		Set("ObjectID", 1).
		Where(goquery.EQ("CommentID", 123)).String()
	require.Nil(t, e)
	assert.Equal(t, "UPDATE `Comment` SET `Content` = 'bar', `ObjectID` = 1 WHERE `CommentID` = 123", sql)
}

func TestUpdate_InvalidField(t *testing.T) {
	sql, e := goquery.Update(&testassets.Comment{}).
		Set("Foo", "bar").
		Set("ObjectID", 1).
		Where(goquery.EQ("CommentID", 123)).String()
	require.NotNil(t, e)
	assert.Equal(t, "UPDATE `Comment` SET `Foo` = 'bar', `ObjectID` = 1 WHERE `CommentID` = 123", sql)
}

func TestDelete(t *testing.T) {
	sql, e := goquery.Delete(&testassets.Comment{}).
		Where(goquery.EQ("CommentID", 123)).String()
	require.Nil(t, e)
	assert.Equal(t, "DELETE FROM `Comment` WHERE `CommentID` = 123", sql)
}

func TestInsert(t *testing.T) {
	sql, e := goquery.Insert(&testassets.Comment{}).
		Set("DateCreated", 1).
		Set("Content", "foo").
		Set("ObjectType", 2).
		Set("ObjectID", 3).
		String()
	require.Nil(t, e)
	assert.Equal(t, "INSERT INTO `Comment` ( `DateCreated`, `Content`, `ObjectType`, `ObjectID` ) VALUES ( 1, 'foo', 2, 3 )", sql)
}

func TestInsert_InvalidFieldName(t *testing.T) {

	sql, e := goquery.Insert(&testassets.Comment{}).
		Set("Foo", "Bar").String()

	require.NotNil(t, e)
	assert.Equal(t, "Invalid Column Name at INSERT...SET in model `Comment` -- Foo", e.Error())
	assert.Equal(t, "INSERT INTO `Comment` ( `Foo` ) VALUES ( Bar )", sql)

}

func TestSelectFields(t *testing.T) {
	sql, e := goquery.Select(&testassets.Job{}).
		Count("JobID", "ProjectsQuoted").
		Sum("TotalPrice", "SalesVolume").
		Sum("GrossProfit", "GM").
		// Field("COALESCE(SUM(TotalPrice), 0)", "SalesVolume").
		// Field("COALESCE(SUM(GrossProfit), 0)", "GM").
		Where(
			goquery.EQ("IsDeleted", 0),
			goquery.And(),
			goquery.Between("AwardDate", 1, 2),
		).
		String()
	require.Nil(t, e)
	assert.Equal(t, "SELECT COUNT(`t`.`JobID`) AS `ProjectsQuoted`, COALESCE(SUM(`t`.`TotalPrice`), 0) AS `SalesVolume`, COALESCE(SUM(`t`.`GrossProfit`), 0) AS `GM` FROM `Job` `t` WHERE `t`.`IsDeleted` = 0 AND `t`.`AwardDate` BETWEEN 1 AND 2", sql)
}

func TestSum_InvalidField(t *testing.T) {
	_, e := goquery.Select(&testassets.Job{}).
		Sum("Foo", "Foo").String()
	require.NotNil(t, e)
	assert.Equal(t, "Invalid Column Name at SELECT...Sum() in model `Job` -- Foo", e.Error())
	// assert.Equal(t, "SELECT SUM(`t`.`Foo`) AS `Foo` FROM `Job` `t`", sql)
}

func TestSelectFields2(t *testing.T) {
	sql, e := goquery.Select(&testassets.Job{}).
		Field("JobID").
		FieldAs("JobID", "foo").
		Where(
			goquery.EQ("IsDeleted", 0),
			goquery.And(),
			goquery.Between("AwardDate", 1, 2),
		).
		String()
	require.Nil(t, e)
	assert.Equal(t, "SELECT `t`.`JobID`, `t`.`JobID` AS `foo` FROM `Job` `t` WHERE `t`.`IsDeleted` = 0 AND `t`.`AwardDate` BETWEEN 1 AND 2", sql)
}

func TestSelectFields3(t *testing.T) {
	sql, e := goquery.Select(&testassets.Job{}).
		Fields(
			goquery.NewField(goquery.FieldTypeBasic, "JobID"),
			goquery.NewField(goquery.FieldTypeBasic, "JobID", "foo"),
		).
		Where(
			goquery.EQ("IsDeleted", 0),
			goquery.And(),
			goquery.Between("AwardDate", 1, 2),
		).
		String()
	require.Nil(t, e)
	assert.Equal(t, "SELECT `t`.`JobID`, `t`.`JobID` AS `foo` FROM `Job` `t` WHERE `t`.`IsDeleted` = 0 AND `t`.`AwardDate` BETWEEN 1 AND 2", sql)
}

func TestSelectAlias(t *testing.T) {
	sql, e := goquery.Select(&testassets.Job{}).
		Alias("j").
		Count("JobID", "ProjectsQuoted").
		// Field("COALESCE(SUM(TotalPrice), 0)", "SalesVolume").
		// Field("COALESCE(SUM(GrossProfit), 0)", "GM").
		Where(
			goquery.EQ("IsDeleted", 0),
			goquery.And(),
			goquery.Between("AwardDate", 1, 2),
		).
		String()
	require.Nil(t, e)
	assert.Equal(t, "SELECT COUNT(`j`.`JobID`) AS `ProjectsQuoted` FROM `Job` `j` WHERE `j`.`IsDeleted` = 0 AND `j`.`AwardDate` BETWEEN 1 AND 2", sql)
}

func TestCountAlias(t *testing.T) {
	sql, e := goquery.Select(&testassets.Job{}).
		Count("JobID", "ProjectsQuoted").
		Where(
			goquery.EQ("IsDeleted", 0),
			goquery.And(),
			goquery.Between("AwardDate", 1, 2),
		).
		Alias("j").
		String()
	require.Nil(t, e)
	assert.Equal(t, "SELECT COUNT(`j`.`JobID`) AS `ProjectsQuoted` FROM `Job` `j` WHERE `j`.`IsDeleted` = 0 AND `j`.`AwardDate` BETWEEN 1 AND 2", sql)
}

func TestSelectExists(t *testing.T) {

	actual, e := goquery.Select(&testassets.Job{}).
		Count("JobID", "ProjectsQuoted").
		Sum("TotalPrice", "SalesVolume").
		Where(
			goquery.Exists(
				goquery.Select(&testassets.JobSales{}).
					Alias("js").
					FieldRaw("1", "n").
					Where(
						goquery.EQF("JobID", "`t`.`JobID`"),
						goquery.And(),
						goquery.EQ("IsDeleted", 0),
						goquery.And(),
						goquery.EQ("UserID", 1),
					),
			),

			//"SELECT 1 FROM `JobSales` `js` WHERE `js`.`JobID` = `j`.`JobID` AND `js`.`IsDeleted` = 0 AND `js`.`UserID` = 1"
			goquery.And(),
			goquery.EQ("IsDeleted", 0),
			goquery.And(),
			goquery.Between("AwardDate", 1, 2),
		).String()

	require.Nil(t, e)
	expected := "SELECT COUNT(`t`.`JobID`) AS `ProjectsQuoted`, COALESCE(SUM(`t`.`TotalPrice`), 0) AS `SalesVolume` FROM `Job` `t` WHERE"
	expected += " EXISTS ( SELECT 1 AS `n` FROM `JobSales` `js` WHERE `js`.`JobID` = `t`.`JobID` AND `js`.`IsDeleted` = 0 AND `js`.`UserID` = 1 )"
	expected += " AND `t`.`IsDeleted` = 0 AND `t`.`AwardDate` BETWEEN 1 AND 2"

	assert.Equal(t, expected, actual)
}

func TestSelectNotExists(t *testing.T) {

	actual, e := goquery.Select(&testassets.Job{}).
		Count("JobID", "ProjectsQuoted").
		Sum("TotalPrice", "SalesVolume").
		Where(
			goquery.NotExists(
				goquery.Select(&testassets.JobSales{}).
					Alias("js").
					FieldRaw("1", "n").
					Where(
						goquery.EQF("JobID", "`t`.`JobID`"),
						goquery.And(),
						goquery.EQ("IsDeleted", 0),
						goquery.And(),
						goquery.EQ("UserID", 1),
					),
			),

			//"SELECT 1 FROM `JobSales` `js` WHERE `js`.`JobID` = `j`.`JobID` AND `js`.`IsDeleted` = 0 AND `js`.`UserID` = 1"
			goquery.And(),
			goquery.EQ("IsDeleted", 0),
			goquery.And(),
			goquery.Between("AwardDate", 1, 2),
		).String()

	require.Nil(t, e)
	expected := "SELECT COUNT(`t`.`JobID`) AS `ProjectsQuoted`, COALESCE(SUM(`t`.`TotalPrice`), 0) AS `SalesVolume` FROM `Job` `t` WHERE"
	expected += " NOT EXISTS ( SELECT 1 AS `n` FROM `JobSales` `js` WHERE `js`.`JobID` = `t`.`JobID` AND `js`.`IsDeleted` = 0 AND `js`.`UserID` = 1 )"
	expected += " AND `t`.`IsDeleted` = 0 AND `t`.`AwardDate` BETWEEN 1 AND 2"

	assert.Equal(t, expected, actual)
}

func TestWhereTypeAll(t *testing.T) {
	q, e := goquery.Select(&testassets.Job{}).Where(goquery.WhereAll()).String()
	require.Nil(t, e)
	expected := "SELECT `t`.* FROM `Job` `t` WHERE 1=1"
	assert.Equal(t, expected, q)

	q, e = goquery.Select(&testassets.Job{}).Where(goquery.WhereAll(), goquery.And(), goquery.EQ("IsDeleted", 0)).String()
	require.Nil(t, e)
	expected = "SELECT `t`.* FROM `Job` `t` WHERE 1=1 AND `t`.`IsDeleted` = 0"
	assert.Equal(t, expected, q)
}

func TestWhere_MultiWheres(t *testing.T) {
	q := goquery.Select(&testassets.Job{}).Where(goquery.WhereAll())
	q.Where(goquery.And(), goquery.EQ("IsDeleted", 0))
	r, e := q.String()

	expected := "SELECT `t`.* FROM `Job` `t` WHERE 1=1 AND `t`.`IsDeleted` = 0"

	assert.Nil(t, e)
	assert.Equal(t, expected, r)
}

// func TestQuerySave_Insert(t *testing.T) {
// 	sql, e := (&testassets.Comment{
// 		Content: null.StringFrom("here is some test content"),
// 	}).Save()
// 	assert.Nil(t, e)
// 	assert.Equal(t, "INSERT INTO `Comment` ( `DateCreated`, `Content`, `ObjectType`, `ObjectID` ) VALUES ( 0, 'here is some test content', 0, 0 )", sql)
// }

// func TestQuerySave_Update(t *testing.T) {
// 	sql, e := Save(&testassets.Comment{
// 		CommentID:  123,
// 		ObjectType: 1,
// 		ObjectID:   2,
// 		Content:    null.StringFrom("here is some test content"),
// 	}).String()
// 	assert.Nil(t, e)
// 	assert.Equal(t, "UPDATE `Comment` SET `Content` = 'here is some test content', `ObjectType` = 1, `ObjectID` = 2 WHERE `CommentID` = 123", sql)
// }

func TestMod(t *testing.T) {
	q, e := goquery.Select(&testassets.Job{}).Where(goquery.Mod("IsDeleted", 1, 0)).String()
	expected := "SELECT `t`.* FROM `Job` `t` WHERE MOD(`t`.`IsDeleted`, 1) = 0"

	assert.Nil(t, e)
	assert.Equal(t, expected, q)
}

func TestModF(t *testing.T) {
	q, e := goquery.Select(&testassets.Job{}).Where(goquery.Modf(1, "IsDeleted", 0)).String()
	expected := "SELECT `t`.* FROM `Job` `t` WHERE MOD(1, `t`.`IsDeleted`) = 0"

	assert.Nil(t, e)
	assert.Equal(t, expected, q)
}

func TestBitAnd(t *testing.T) {
	q, e := goquery.Select(&testassets.Job{}).Where(goquery.BitAnd("IsDeleted", 1, 0)).String()
	expected := "SELECT `t`.* FROM `Job` `t` WHERE `t`.`IsDeleted` & 1 = 0"

	assert.Nil(t, e)
	assert.Equal(t, expected, q)
}

func TestModAndBitwise(t *testing.T) {
	// now := time.Now()

	dayBW := int64(16)       // int64(math.Pow(2, float64(now.Day())))
	weekdayBW := int64(1)    // int64(math.Pow(2, float64(now.Weekday())))
	hourBW := int64(2048)    // int64(math.Pow(2, float64(now.Hour())))
	minute := int64(56)      // int64(time.Now().Minute())
	seconds := 1625414209657 // time.Now().UnixNano() / 1000000
	// SELECT * FROM TaskBatchSchedule tbs WHERE tbs.IsActive = 1 AND tbs.IsDeleted = 0 AND ( tbs.DOM = 0 OR 16 & tbs.DOM = 16 ) AND ( tbs.DOW = 0 OR 1 & tbs.DOW = 1 ) AND ( tbs.HOD = 0 OR 2048 & tbs.HOD = 2048 ) AND ( MOD(56, tbs.MOH ) = 0 AND ( tbs.LastRunDate + 60000 < 1625414209657 )

	q := "SELECT `t`.* FROM `TaskBatchSchedule` `t` WHERE `t`.`IsActive` = 1 AND `t`.`IsDeleted` = 0 "
	q += fmt.Sprintf("AND ( `t`.`DOM` = 0 OR `t`.`DOM` & %d = %d ) ", dayBW, dayBW)
	q += fmt.Sprintf("AND ( `t`.`DOW` = 0 OR `t`.`DOW` & %d = %d ) ", weekdayBW, weekdayBW)
	q += fmt.Sprintf("AND ( `t`.`HOD` = 0 OR `t`.`HOD` & %d = %d ) ", hourBW, hourBW)
	// Every minute of the hour
	q += fmt.Sprintf("AND ( MOD(%d, `t`.`MOH`) = 0 ) ", minute)
	// It has been atleast 60 seconds since the last run
	q += fmt.Sprintf("AND ( `t`.`LastRunDate` + 60000 < %d )", seconds)

	actual, e := goquery.Select(&testassets.TaskBatchSchedule{}).
		Where(
			goquery.EQ("IsActive", 1),
			goquery.And(),
			goquery.EQ("IsDeleted", 0),
			goquery.And(goquery.EQ("DOM", 0), goquery.Or(), goquery.BitAnd("DOM", dayBW, dayBW)),
			goquery.And(goquery.EQ("DOW", 0), goquery.Or(), goquery.BitAnd("DOW", weekdayBW, weekdayBW)),
			goquery.And(goquery.EQ("HOD", 0), goquery.Or(), goquery.BitAnd("HOD", hourBW, hourBW)),
			goquery.And(goquery.Modf(minute, "MOH", 0)),
			goquery.And(goquery.Rawf("`t`.`LastRunDate` + 60000 < %d", seconds)),
		).
		String()

	// fmt.Println(q)

	assert.Nil(t, e)
	assert.Equal(t, q, actual)

}

func TestGrouping(t *testing.T) {
	expected := "SELECT `t`.* FROM `FiscalYear` `t` WHERE `t`.`IsDeleted` = 0 AND ( ( `t`.`DateFrom` BETWEEN 1 AND 2 ) OR ( `t`.`DateTo` BETWEEN 3 AND 4 ) )"
	actual, e := goquery.Select(&testassets.FiscalYear{}).
		Where(
			goquery.EQ("IsDeleted", 0),
			goquery.And(
				goquery.Paren(goquery.Between("DateFrom", 1, 2)),
				goquery.Or(),
				goquery.Paren(goquery.Between("DateTo", 3, 4)),
			),
		).
		String()

	assert.Nil(t, e)
	assert.Equal(t, expected, actual)
}

func TestRaw(t *testing.T) {

	q, e := goquery.Raw(&testassets.TaskBatchSchedule{}, "SELECT * FROM `TaskBatchSchedule` WHERE 1=1 ORDER BY `t`.`LastRunDate` DESC LIMIT 1 OFFSET 2").
		String()

	assert.Nil(t, e)
	assert.Equal(t, "SELECT * FROM `TaskBatchSchedule` WHERE 1=1 ORDER BY `t`.`LastRunDate` DESC LIMIT 1 OFFSET 2", q)
}

func TestAnds(t *testing.T) {
	q1, e1 := goquery.Select(&testassets.FiscalYear{}).
		Where(
			goquery.Ands(
				goquery.EQ(testassets.FiscalYear_Column_IsDeleted, 0),
				goquery.Between(testassets.FiscalYear_Column_DateFrom, 1, 2),
				goquery.EQ(testassets.FiscalYear_Column_IsLocked, 0),
				goquery.EQ(testassets.FiscalYear_Column_Year, 2021),
			),
		).String()

	q2, e2 := goquery.Select(&testassets.FiscalYear{}).
		Where(
			goquery.Ands(
				goquery.EQ(testassets.FiscalYear_Column_IsDeleted, 0),
				goquery.Between(testassets.FiscalYear_Column_DateFrom, 1, 2),
				goquery.EQ(testassets.FiscalYear_Column_IsLocked, 0),
				goquery.EQ(testassets.FiscalYear_Column_Year, 2021),
				nil,
			),
		).String()

	q3, e3 := goquery.Select(&testassets.FiscalYear{}).
		Where(
			goquery.Ands(
				goquery.EQ(testassets.FiscalYear_Column_IsDeleted, 0),
				goquery.Between(testassets.FiscalYear_Column_DateFrom, 1, 2),
				goquery.EQ(testassets.FiscalYear_Column_IsLocked, 0),
				nil,
				goquery.EQ(testassets.FiscalYear_Column_Year, 2021),
			),
		).String()
	q4, e4 := goquery.Select(&testassets.FiscalYear{}).
		Where(
			goquery.Ands(
				nil,
				goquery.EQ(testassets.FiscalYear_Column_IsDeleted, 0),
				goquery.Between(testassets.FiscalYear_Column_DateFrom, 1, 2),
				goquery.EQ(testassets.FiscalYear_Column_IsLocked, 0),
				goquery.EQ(testassets.FiscalYear_Column_Year, 2021),
			),
		).String()

	assert.Nil(t, e1)
	assert.Equal(t, "SELECT `t`.* FROM `FiscalYear` `t` WHERE `t`.`IsDeleted` = 0 AND `t`.`DateFrom` BETWEEN 1 AND 2 AND `t`.`IsLocked` = 0 AND `t`.`Year` = 2021", q1)
	assert.Nil(t, e2)
	assert.Equal(t, "SELECT `t`.* FROM `FiscalYear` `t` WHERE `t`.`IsDeleted` = 0 AND `t`.`DateFrom` BETWEEN 1 AND 2 AND `t`.`IsLocked` = 0 AND `t`.`Year` = 2021", q2)
	assert.Nil(t, e3)
	assert.Equal(t, "SELECT `t`.* FROM `FiscalYear` `t` WHERE `t`.`IsDeleted` = 0 AND `t`.`DateFrom` BETWEEN 1 AND 2 AND `t`.`IsLocked` = 0 AND `t`.`Year` = 2021", q3)
	assert.Nil(t, e4)
	assert.Equal(t, "SELECT `t`.* FROM `FiscalYear` `t` WHERE `t`.`IsDeleted` = 0 AND `t`.`DateFrom` BETWEEN 1 AND 2 AND `t`.`IsLocked` = 0 AND `t`.`Year` = 2021", q4)
}

func TestOrs(t *testing.T) {
	q1, e1 := goquery.Select(&testassets.FiscalYear{}).
		Where(
			goquery.Ors(
				goquery.EQ(testassets.FiscalYear_Column_IsDeleted, 0),
				goquery.EQ(testassets.FiscalYear_Column_IsLocked, 0),
				goquery.EQ(testassets.FiscalYear_Column_Year, 2021),
			),
		).String()

	q2, e2 := goquery.Select(&testassets.FiscalYear{}).
		Where(
			goquery.Ors(
				goquery.EQ(testassets.FiscalYear_Column_IsDeleted, 0),
				goquery.EQ(testassets.FiscalYear_Column_IsLocked, 0),
				goquery.EQ(testassets.FiscalYear_Column_Year, 2021),
				nil,
			),
		).String()

	q3, e3 := goquery.Select(&testassets.FiscalYear{}).
		Where(
			goquery.Ors(
				goquery.EQ(testassets.FiscalYear_Column_IsDeleted, 0),
				goquery.EQ(testassets.FiscalYear_Column_IsLocked, 0),
				nil,
				goquery.EQ(testassets.FiscalYear_Column_Year, 2021),
			),
		).String()

	q4, e4 := goquery.Select(&testassets.FiscalYear{}).
		Where(
			goquery.Ors(
				nil,
				goquery.EQ(testassets.FiscalYear_Column_IsDeleted, 0),
				goquery.EQ(testassets.FiscalYear_Column_IsLocked, 0),
				goquery.EQ(testassets.FiscalYear_Column_Year, 2021),
			),
		).String()

	assert.Nil(t, e1)
	assert.Equal(t, "SELECT `t`.* FROM `FiscalYear` `t` WHERE `t`.`IsDeleted` = 0 OR `t`.`IsLocked` = 0 OR `t`.`Year` = 2021", q1)
	assert.Nil(t, e2)
	assert.Equal(t, "SELECT `t`.* FROM `FiscalYear` `t` WHERE `t`.`IsDeleted` = 0 OR `t`.`IsLocked` = 0 OR `t`.`Year` = 2021", q2)
	assert.Nil(t, e3)
	assert.Equal(t, "SELECT `t`.* FROM `FiscalYear` `t` WHERE `t`.`IsDeleted` = 0 OR `t`.`IsLocked` = 0 OR `t`.`Year` = 2021", q3)
	assert.Nil(t, e4)
	assert.Equal(t, "SELECT `t`.* FROM `FiscalYear` `t` WHERE `t`.`IsDeleted` = 0 OR `t`.`IsLocked` = 0 OR `t`.`Year` = 2021", q4)
}

// func TestWhereSelect(t *testing.T) {
// 	expected := "SELECT `t`.* FROM `Comment` `t` WHERE `t`.`CommentID` < (SELECT QuoteNumberFullInt FROM QuoteNumber WHERE QuoteNumberID = %d) AND `t`.`IsDeleted` = 0 AND `t`.`JobID` > 0 ORDER BY `t`.`QuoteNumberFullInt` DESC LIMIT 1"

// 	q, e := goquery.Select(&testassets.Comment{}).Where(
// 		goquery.LT(testassets.Comment_Column_CommentID, )
// 	).String()

// 	assert.Nil(t, e)
// 	assert.Equal(t, expected, q)
// }

func TestEscapeString(t *testing.T) {

	result := goquery.EscapeString("I'm a string")
	assert.Equal(t, `I\'m a string`, result)

	result = goquery.EscapeString(`I"m a string`)
	assert.Equal(t, `I\"m a string`, result)

}

func TestMin(t *testing.T) {
	q1, e1 := goquery.Select(&testassets.FiscalYear{}).Min(testassets.FiscalYear_Column_Year, "MinYear").Where(goquery.EQ(testassets.FiscalYear_Column_IsDeleted, 0)).String()

	assert.Nil(t, e1)
	assert.Equal(t, "SELECT COALESCE(MIN(`t`.`Year`), 0) AS `MinYear` FROM `FiscalYear` `t` WHERE `t`.`IsDeleted` = 0", q1)
}

func TestMax(t *testing.T) {
	q1, e1 := goquery.Select(&testassets.FiscalYear{}).Max(testassets.FiscalYear_Column_Year, "MaxYear").Where(goquery.EQ(testassets.FiscalYear_Column_IsDeleted, 0)).String()

	assert.Nil(t, e1)
	assert.Equal(t, "SELECT COALESCE(MAX(`t`.`Year`), 0) AS `MaxYear` FROM `FiscalYear` `t` WHERE `t`.`IsDeleted` = 0", q1)
}

func ExampleMod() {
	var modString = goquery.Mod("foo", 2, 1)
	var result, _ = goquery.Select(&testassets.Comment{}).
		Where(modString).
		String()
	fmt.Println(result)

	// Output: SELECT `t`.* FROM `Comment` `t` WHERE MOD(`t`.`foo`, 2) = 1
}

func TestAvgAndCountsAndSums(t *testing.T) {

	q, e := goquery.Select(&testassets.FiscalYear{}).
		Count(testassets.FiscalYear_Column_FiscalYearID, "Overall").
		Sum(testassets.FiscalYear_Column_FiscalYearID, "ChangesFound").
		Avg(testassets.FiscalYear_Column_FiscalYearID, "AverageChangesFound").String()

	assert.Nil(t, e)
	assert.Equal(t, "SELECT COUNT(`t`.`FiscalYearID`) AS `Overall`, COALESCE(SUM(`t`.`FiscalYearID`), 0) AS `ChangesFound`, COALESCE(AVG(`t`.`FiscalYearID`), 0) AS `AverageChangesFound` FROM `FiscalYear` `t`", q)

}
