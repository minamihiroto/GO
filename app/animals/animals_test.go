package animals

import(
	"testing"
)

//Testで始まる関数がテストの単位を表す
//animalsパッケージに定義されている3つの関数に対応したテスト用関数を定義
//今回描いたのはユニットテストのようなもの
func TestElephantFeed(t *testing.T){
	expect := "Grass"//期待
	actual := ElephantFeed()//実際

	//期待と実際が違うならエラーを出す
	//Errorfはエラーを出す関数
	if expect != actual{
		t.Errorf("期待：%s,実際：%s",expect,actual)//%sは文字列またはスライスそのまま
	}
}

func TestMonkeyFeed(t *testing.T){
	expect := "Banana"
	actual := MonkeyFeed()

	if expect != actual{
		t.Errorf("%s != %s",expect,actual)
	}
}

func TestRabbitFeed(t *testing.T){
	expect := "Carrot"
	actual := RabbitFeed()

	if expect != actual{
		t.Errorf("%s != %s",expect,actual)
	}
}

//$ go test gomodtest/animals でテストを実行
//上記コマンドからもわかる通りテストはディレクトリ単位で行える