#交易(Transactions)

**交易(Transactions)**是比特币的核心，而区块链的唯一的目的，也是为了能够按期可靠地存储交易。
在区块链中，交易一旦被创建，就没有任何人能够再去修改或时删除它。

## 比特币交易

比特币交易由一些输入(input)和输出(output)组合而来：

```go
type Transaction struct {
	ID 		[]byte
	Vin 	[]TXInput
	Vout 	[]TXOutput
}
```

每一笔新的交的输入会引用(reference)之前一笔交易的输出(这里有个例外，就是coinbase交易，它是奖励机制中对于打包区块节点的奖励)。


