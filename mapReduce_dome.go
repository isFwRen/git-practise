package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/mr"
	"time"
)

// Token 输入数据类型
type Token struct {
	Id      int64
	ChainId int64
}

// WalletToken Map 阶段输出数据类型
type WalletToken struct {
	WalletId int64
	TokenId  int64
	ChainId  int64
	IsHome   int
	Status   int
	CreateAt time.Time
	IsCommon int
}

// WalletTokens Reduce 阶段输出数据类型
type WalletTokens []*WalletToken

func main() {
	// 模拟输入数据
	commonTokenList := []*Token{
		{Id: 1, ChainId: 100},
		{Id: 2, ChainId: 200},
		{Id: 3, ChainId: 300},
	}

	// 定义 WalletId
	walletId := int64(1)

	// 使用 MapReduce 函数
	tokenList, err := mr.MapReduce[*Token, *WalletToken, WalletTokens](func(source chan<- *Token) {
		for _, t := range commonTokenList {
			source <- t
		}
	}, func(item *Token, writer mr.Writer[*WalletToken], cancel func(error)) {
		wt := &WalletToken{
			WalletId: walletId,
			TokenId:  item.Id,
			ChainId:  item.ChainId,
			IsHome:   0,
			Status:   0,
			CreateAt: time.Now().UTC(),
			IsCommon: 1,
		}
		fmt.Println("wt：", *wt)
		writer.Write(wt)
	}, func(pipe <-chan *WalletToken, writer mr.Writer[WalletTokens], cancel func(error)) {
		var res WalletTokens
		for t := range pipe {
			res = append(res, t)
		}
		writer.Write(res)
	})

	if err != nil {
		fmt.Printf("MapReduce error: %v\n", err)
		return
	}

	// 输出结果
	for _, token := range tokenList {
		fmt.Printf("WalletId: %d, TokenId: %d, ChainId: %d\n", token.WalletId, token.TokenId, token.ChainId)
	}
}
