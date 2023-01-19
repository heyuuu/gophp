// <<generate>>

package standard

/**
 * Querybuf
 */
type Querybuf struct /* union */ {
	qb1 HEADER
	qb2 []u_char
}

func (this Querybuf) GetQb1() HEADER         { return this.qb1 }
func (this *Querybuf) SetQb1(value HEADER)   { this.qb1 = value }
func (this Querybuf) GetQb2() []u_char       { return this.qb2 }
func (this *Querybuf) SetQb2(value []u_char) { this.qb2 = value }
