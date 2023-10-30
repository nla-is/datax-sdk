package main

// #include <stdint.h>
// typedef const char * const_char_p;
// typedef const unsigned char * const_unsigned_char_p;
// typedef uintptr_t datax_sdk_v2_message;
import "C"

func main() {
}

// datax_sdk_v2_initialize Initialize DataX SDK
//
//export datax_sdk_v2_initialize
func datax_sdk_v2_initialize() {
	panic("datax_sdk_v2_begin_initialize not implemented")
}

// datax_sdk_v2_next Receives an input message.
//
// The input message must be freed using the function datax_sdk_v2_message_close()
//
//export datax_sdk_v2_next
func datax_sdk_v2_next() C.datax_sdk_v2_message {
	panic("datax_sdk_v2_next not implemented")
}

// datax_sdk_v2_emit Publishes a message
//
// The message data should be encoded in CBOR format (http://cbor.io/)
//
//export datax_sdk_v2_emit
func datax_sdk_v2_emit(data C.const_unsigned_char_p, data_size C.int32_t, reference C.const_char_p) {
	panic("datax_sdk_v2_emit not implemented")
}

// datax_sdk_v2_message_close Frees the resources of the message
//
//export datax_sdk_v2_message_close
func datax_sdk_v2_message_close(message C.datax_sdk_v2_message) {
	panic("datax_sdk_v2_message_close not implemented")
}

// datax_sdk_v2_message_reference Obtains the reference of the message
//
//export datax_sdk_v2_message_reference
func datax_sdk_v2_message_reference(message C.datax_sdk_v2_message) C.const_char_p {
	panic("datax_sdk_v2_message_reference not implemented")
}

// datax_sdk_v2_message_stream Obtains the name of the stream that generated the message
//
//export datax_sdk_v2_message_stream
func datax_sdk_v2_message_stream(message C.datax_sdk_v2_message) C.const_char_p {
	panic("datax_sdk_v2_message_stream not implemented")
}

// datax_sdk_v2_message_data Obtains the data of the message
//
// The data is encoded in CBOR format (http://cbor.io/)
//
//export datax_sdk_v2_message_data
func datax_sdk_v2_message_data(message C.datax_sdk_v2_message) C.const_unsigned_char_p {
	panic("datax_sdk_v2_message_data not implemented")
}

// datax_sdk_v2_message_data_size Obtains the size of the data in the message
//
//export datax_sdk_v2_message_data_size
func datax_sdk_v2_message_data_size(message C.datax_sdk_v2_message) C.int32_t {
	panic("datax_sdk_v2_message_data_size not implemented")
}
