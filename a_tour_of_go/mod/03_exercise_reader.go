package main

type MyReader struct {
}

type MyReaderError bool

func (MyReaderError) Error() string {
	return "too short b capacity"
}

func (MyReader) Read(b []byte) (int, error) {
	if cap(b) < 1 {
		return 0, MyReaderError(true)
	}
	b[0] = 'A'
	return 1, nil
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
// func (r *MyReader) Read(b []byte) (int, error) {
// 	if len(r.str) <= r.done {
// 		return 0, io.EOF
// 	}
// 	s := r.str[r.done:len(b)]
// 	ln := len(s)
// 	r.done = r.done + ln
// 	return ln, nil
// }
