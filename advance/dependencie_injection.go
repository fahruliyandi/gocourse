package advance

//==========DEPENDENCIES INJECTION============
// type kamera interface {
// 	jepret() string
// }

// type cannonKamera struct{}
// type hapeKamera struct{}

// func (ck cannonKamera) jepret() string {
// 	return "Jepret dari cannon"
// }

// func (hk hapeKamera) jepret() string {
// 	return "jepret dari hape"
// }

// type tukangPoto struct {
// 	km kamera
// }

// func (tp tukangPoto) tp_jepret() {
// 	fmt.Println(tp.km.jepret())
// }

// func newTukangPoto(k kamera) tukangPoto {
// 	return tukangPoto{km: k}
// }

// func main() {

// 	cannon := cannonKamera{}
// 	hp := hapeKamera{}
// 	tp := newTukangPoto(cannon)
// 	tp.tp_jepret()
// 	tp = newTukangPoto(hp)
// 	tp.tp_jepret()

// }

//=========implementasi interface========
// type kendaraan interface {
// 	jalan() string
// }

// type mobil struct{}
// type motor struct{}
// type skuter struct{}

// func (mb mobil) jalan() string {
// 	return "mobil jalan"
// }

// func (mt motor) jalan() string {
// 	return "motor jalan"
// }

// func (sk skuter) jalan() string {
// 	return "skuter jalan"
// }

// func prosesJalan(k kendaraan) {
// 	fmt.Println(k.jalan())
// }

// func main() {

// 	mb := mobil{}
// 	mt := motor{}
// 	sk := skuter{}

// 	prosesJalan(mb)
// 	prosesJalan(mt)
// 	prosesJalan(sk)

// }
