package monster

import "testing"

//测试用例，测试store方法
func TestScore(t *testing.T){

	//先创建一个Monster的实例
	monster := Monster{
		Name:"红孩儿",
		Age:23,
		Skill:"三位真火",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store()错误，希望为%v,结果为%v",true,res)
	}
	t.Logf("monster.Store()测试成功")
}

func TestRescore(t *testing.T){

	var monster = &Monster{}
	res := monster.Restore()
	if !res {
		t.Fatalf("monster.Restore()错误，希望为%v,结果为%v",true,res)
	}
	//进一步判断
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.Restore()错误，希望为%v,结果为%v","红孩儿",monster.Name)
	}
	t.Logf("monster.Restore()测试成功")
}

func main () {

	
}

