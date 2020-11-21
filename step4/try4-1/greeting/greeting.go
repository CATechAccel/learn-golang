package greeting

//外部から利用したい場合は、関数名などは大文字で始める
func Do(x *string) {
	*x = "こんにちは"
	return
}
