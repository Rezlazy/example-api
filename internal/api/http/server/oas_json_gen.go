// Code generated by ogen, DO NOT EDIT.

package server

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *AuthLoginRequestSchema) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *AuthLoginRequestSchema) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("email")
		e.Str(s.Email)
	}
	{
		e.FieldStart("password")
		e.Str(s.Password)
	}
}

var jsonFieldsNameOfAuthLoginRequestSchema = [2]string{
	0: "email",
	1: "password",
}

// Decode decodes AuthLoginRequestSchema from json.
func (s *AuthLoginRequestSchema) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AuthLoginRequestSchema to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "email":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Email = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"email\"")
			}
		case "password":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Password = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"password\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode AuthLoginRequestSchema")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfAuthLoginRequestSchema) {
					name = jsonFieldsNameOfAuthLoginRequestSchema[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *AuthLoginRequestSchema) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AuthLoginRequestSchema) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ErrorResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ErrorResponse) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("error_message")
		e.Str(s.ErrorMessage)
	}
}

var jsonFieldsNameOfErrorResponse = [1]string{
	0: "error_message",
}

// Decode decodes ErrorResponse from json.
func (s *ErrorResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ErrorResponse to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "error_message":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.ErrorMessage = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"error_message\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ErrorResponse")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfErrorResponse) {
					name = jsonFieldsNameOfErrorResponse[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ErrorResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ErrorResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Item) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Item) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("title")
		e.Str(s.Title)
	}
	{
		if s.Description.Set {
			e.FieldStart("description")
			s.Description.Encode(e)
		}
	}
	{
		e.FieldStart("list_id")
		e.Int64(s.ListID)
	}
	{
		e.FieldStart("id")
		e.Int64(s.ID)
	}
}

var jsonFieldsNameOfItem = [4]string{
	0: "title",
	1: "description",
	2: "list_id",
	3: "id",
}

// Decode decodes Item from json.
func (s *Item) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Item to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "title":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Title = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "description":
			if err := func() error {
				s.Description.Reset()
				if err := s.Description.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"description\"")
			}
		case "list_id":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Int64()
				s.ListID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"list_id\"")
			}
		case "id":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Int64()
				s.ID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Item")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00001101,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfItem) {
					name = jsonFieldsNameOfItem[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Item) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Item) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ItemCommonData) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ItemCommonData) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("title")
		e.Str(s.Title)
	}
	{
		if s.Description.Set {
			e.FieldStart("description")
			s.Description.Encode(e)
		}
	}
}

var jsonFieldsNameOfItemCommonData = [2]string{
	0: "title",
	1: "description",
}

// Decode decodes ItemCommonData from json.
func (s *ItemCommonData) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ItemCommonData to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "title":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Title = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "description":
			if err := func() error {
				s.Description.Reset()
				if err := s.Description.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"description\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ItemCommonData")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfItemCommonData) {
					name = jsonFieldsNameOfItemCommonData[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ItemCommonData) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ItemCommonData) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ItemCreateInput) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ItemCreateInput) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("title")
		e.Str(s.Title)
	}
	{
		if s.Description.Set {
			e.FieldStart("description")
			s.Description.Encode(e)
		}
	}
	{
		e.FieldStart("list_id")
		e.Int64(s.ListID)
	}
}

var jsonFieldsNameOfItemCreateInput = [3]string{
	0: "title",
	1: "description",
	2: "list_id",
}

// Decode decodes ItemCreateInput from json.
func (s *ItemCreateInput) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ItemCreateInput to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "title":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Title = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "description":
			if err := func() error {
				s.Description.Reset()
				if err := s.Description.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"description\"")
			}
		case "list_id":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Int64()
				s.ListID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"list_id\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ItemCreateInput")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000101,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfItemCreateInput) {
					name = jsonFieldsNameOfItemCreateInput[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ItemCreateInput) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ItemCreateInput) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *List) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *List) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		e.Int64(s.ID)
	}
	{
		e.FieldStart("name")
		e.Str(s.Name)
	}
	{
		if s.Items != nil {
			e.FieldStart("items")
			e.ArrStart()
			for _, elem := range s.Items {
				elem.Encode(e)
			}
			e.ArrEnd()
		}
	}
}

var jsonFieldsNameOfList = [3]string{
	0: "id",
	1: "name",
	2: "items",
}

// Decode decodes List from json.
func (s *List) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode List to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int64()
				s.ID = int64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "name":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "items":
			if err := func() error {
				s.Items = make([]Item, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Item
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Items = append(s.Items, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"items\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode List")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfList) {
					name = jsonFieldsNameOfList[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *List) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *List) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ListCreateInput) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ListCreateInput) encodeFields(e *jx.Encoder) {
	{
		if s.Items != nil {
			e.FieldStart("items")
			e.ArrStart()
			for _, elem := range s.Items {
				elem.Encode(e)
			}
			e.ArrEnd()
		}
	}
	{
		e.FieldStart("name")
		e.Str(s.Name)
	}
}

var jsonFieldsNameOfListCreateInput = [2]string{
	0: "items",
	1: "name",
}

// Decode decodes ListCreateInput from json.
func (s *ListCreateInput) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ListCreateInput to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "items":
			if err := func() error {
				s.Items = make([]ItemCommonData, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem ItemCommonData
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Items = append(s.Items, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"items\"")
			}
		case "name":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ListCreateInput")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000010,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfListCreateInput) {
					name = jsonFieldsNameOfListCreateInput[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ListCreateInput) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ListCreateInput) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ListListResponseSchema as json.
func (s ListListResponseSchema) Encode(e *jx.Encoder) {
	unwrapped := []List(s)

	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes ListListResponseSchema from json.
func (s *ListListResponseSchema) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ListListResponseSchema to nil")
	}
	var unwrapped []List
	if err := func() error {
		unwrapped = make([]List, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem List
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = ListListResponseSchema(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s ListListResponseSchema) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ListListResponseSchema) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Token) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Token) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("access_token")
		e.Str(s.AccessToken)
	}
}

var jsonFieldsNameOfToken = [1]string{
	0: "access_token",
}

// Decode decodes Token from json.
func (s *Token) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Token to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "access_token":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.AccessToken = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"access_token\"")
			}
		default:
			return errors.Errorf("unexpected field %q", k)
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Token")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfToken) {
					name = jsonFieldsNameOfToken[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Token) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Token) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
