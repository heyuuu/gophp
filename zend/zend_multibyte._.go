package zend

type ZendEncoding string

type ZendEncodingFilter func(str **uint8, str_length *int, buf *uint8, length int) int
