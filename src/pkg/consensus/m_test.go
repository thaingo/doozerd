package consensus

import (
	"fmt"
	"goprotobuf.googlecode.com/hg/proto"
)

func (x *msg_Cmd) Format(f fmt.State, c int) {
	if c == 'v' && f.Flag('#') && x != nil {
		fmt.Fprintf(f, "msg_%s", msg_Cmd_name[int32(*x)])
		return
	}

	s := "%"
	for i := 0; i < 128; i++ {
		if f.Flag(i) {
			s += string(i)
		}
	}
	if w, ok := f.Width(); ok {
		s += fmt.Sprintf("%d", w)
	}
	if p, ok := f.Precision(); ok {
		s += fmt.Sprintf(".%d", p)
	}
	s += string(c)
	fmt.Fprintf(f, s, (*int32)(x))
}

// For testing convenience
func newVote(i int64, vval string) *msg {
	return &msg{Cmd: vote, Vrnd: &i, Value: []byte(vval)}
}

// For testing convenience
func newVoteFrom(from string, i int64, vval string) packet {
	m := newVote(i, vval)
	m.Seqn = proto.Int64(1)
	return packet{from, *m}
}

// For testing convenience
func newNominate(crnd int64, v string) *msg {
	return &msg{Cmd: nominate, Crnd: &crnd, Value: []byte(v)}
}

// For testing convenience
func newNominateSeqn1(crnd int64, v string) *msg {
	m := newNominate(crnd, v)
	m.Seqn = proto.Int64(1)
	return m
}

// For testing convenience
func newRsvp(i, vrnd int64, vval string) *msg {
	return &msg{
		Cmd:   rsvp,
		Crnd:  &i,
		Vrnd:  &vrnd,
		Value: []byte(vval),
	}
}

// For testing convenience
func newRsvpFrom(from string, i, vrnd int64, vval string) packet {
	m := newRsvp(i, vrnd, vval)
	m.Seqn = proto.Int64(1)
	return packet{from, *m}
}

// For testing convenience
func newInvite(crnd int64) *msg {
	return &msg{Cmd: invite, Crnd: &crnd}
}

// For testing convenience
func newInviteSeqn1(rnd int64) *msg {
	m := newInvite(rnd)
	m.Seqn = proto.Int64(1)
	return m
}

// For testing convenience
func newPropose(val string) *msg {
	return &msg{Cmd: propose, Value: []byte(val)}
}

// For testing convenience
func newLearn(val string) *msg {
	return &msg{Cmd: learn, Value: []byte(val)}
}
