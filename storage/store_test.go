/// @Author winjeg,  winjeg@qq.com
/// All rights reserved to winjeg

package storage

import "testing"

func TestParseSql(t *testing.T) {
	ParseSql("SELECT cola AS a, colb AS b FROM sales.abc aa  LEFT JOIN def bb ON aa.cola = bb.cola LEFT JOIN cc ON cc.id = aa.id WHERE aa.id > 10 ")
}
