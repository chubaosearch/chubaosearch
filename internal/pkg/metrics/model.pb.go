// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.0
// source: model.proto

package metrics

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MetricType describes the type of metric.
type MetricType int32

const (
	// UNSET expresses that the metric's type wasn't explicitly set.
	MetricType_UNTYPED     MetricType = 0
	MetricType_COUNTER     MetricType = 1
	MetricType_COUNTERRATE MetricType = 2
	MetricType_GAUGE       MetricType = 3
	MetricType_HISTOGRAM   MetricType = 4
)

// Enum value maps for MetricType.
var (
	MetricType_name = map[int32]string{
		0: "UNTYPED",
		1: "COUNTER",
		2: "COUNTERRATE",
		3: "GAUGE",
		4: "HISTOGRAM",
	}
	MetricType_value = map[string]int32{
		"UNTYPED":     0,
		"COUNTER":     1,
		"COUNTERRATE": 2,
		"GAUGE":       3,
		"HISTOGRAM":   4,
	}
)

func (x MetricType) Enum() *MetricType {
	p := new(MetricType)
	*p = x
	return p
}

func (x MetricType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricType) Descriptor() protoreflect.EnumDescriptor {
	return file_model_proto_enumTypes[0].Descriptor()
}

func (MetricType) Type() protoreflect.EnumType {
	return &file_model_proto_enumTypes[0]
}

func (x MetricType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricType.Descriptor instead.
func (MetricType) EnumDescriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{0}
}

// Unit describes how the metric's units should be displayed.
type Unit int32

const (
	// UNSET expresses that the metric's DisplayUnit wasn't explicitly set.
	Unit_UNSET Unit = 0
	// BYTES expresses that the metric's measurement is in bytes.
	Unit_BYTES Unit = 1
	// CONST expresses that the metric's measurement is a constant value.
	Unit_CONST Unit = 2
	// COUNT expresses that the metric's measurement is a count.
	Unit_COUNT Unit = 3
	// NANOSECONDS expresses that the metric's measurement is in nanoseconds.
	Unit_NANOSECONDS Unit = 4
	// PERCENT expresses that the metric's measurement is a percentage value.
	Unit_PERCENT Unit = 5
	// SECONDS expresses that the metric's measurement is in seconds.
	Unit_SECONDS Unit = 6
)

// Enum value maps for Unit.
var (
	Unit_name = map[int32]string{
		0: "UNSET",
		1: "BYTES",
		2: "CONST",
		3: "COUNT",
		4: "NANOSECONDS",
		5: "PERCENT",
		6: "SECONDS",
	}
	Unit_value = map[string]int32{
		"UNSET":       0,
		"BYTES":       1,
		"CONST":       2,
		"COUNT":       3,
		"NANOSECONDS": 4,
		"PERCENT":     5,
		"SECONDS":     6,
	}
)

func (x Unit) Enum() *Unit {
	p := new(Unit)
	*p = x
	return p
}

func (x Unit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Unit) Descriptor() protoreflect.EnumDescriptor {
	return file_model_proto_enumTypes[1].Descriptor()
}

func (Unit) Type() protoreflect.EnumType {
	return &file_model_proto_enumTypes[1]
}

func (x Unit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Unit.Descriptor instead.
func (Unit) EnumDescriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{1}
}

type LabelPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *LabelPair) Reset() {
	*x = LabelPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelPair) ProtoMessage() {}

func (x *LabelPair) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelPair.ProtoReflect.Descriptor instead.
func (*LabelPair) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{0}
}

func (x *LabelPair) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LabelPair) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Metadata holds metadata about a metric.
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Help   string       `protobuf:"bytes,2,opt,name=help,proto3" json:"help,omitempty"`
	Unit   Unit         `protobuf:"varint,3,opt,name=unit,proto3,enum=metrics.Unit" json:"unit,omitempty"`
	Labels []*LabelPair `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{1}
}

func (x *Metadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metadata) GetHelp() string {
	if x != nil {
		return x.Help
	}
	return ""
}

func (x *Metadata) GetUnit() Unit {
	if x != nil {
		return x.Unit
	}
	return Unit_UNSET
}

func (x *Metadata) GetLabels() []*LabelPair {
	if x != nil {
		return x.Labels
	}
	return nil
}

// MetricFamily contains a list of metrics with the same name (a metric with multiple labels).
type MetricFamily struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Help    string        `protobuf:"bytes,2,opt,name=help,proto3" json:"help,omitempty"`
	Type    MetricType    `protobuf:"varint,3,opt,name=type,proto3,enum=metrics.MetricType" json:"type,omitempty"`
	Metrics []*MetricData `protobuf:"bytes,4,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *MetricFamily) Reset() {
	*x = MetricFamily{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricFamily) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricFamily) ProtoMessage() {}

func (x *MetricFamily) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricFamily.ProtoReflect.Descriptor instead.
func (*MetricFamily) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{2}
}

func (x *MetricFamily) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetricFamily) GetHelp() string {
	if x != nil {
		return x.Help
	}
	return ""
}

func (x *MetricFamily) GetType() MetricType {
	if x != nil {
		return x.Type
	}
	return MetricType_UNTYPED
}

func (x *MetricFamily) GetMetrics() []*MetricData {
	if x != nil {
		return x.Metrics
	}
	return nil
}

// MetricData holds exported data of metrics
type MetricData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Counter     *CounterData     `protobuf:"bytes,1,opt,name=counter,proto3" json:"counter,omitempty"`
	CounterRate *CounterRateData `protobuf:"bytes,2,opt,name=counter_rate,json=counterRate,proto3" json:"counter_rate,omitempty"`
	Gauge       *GaugeData       `protobuf:"bytes,3,opt,name=gauge,proto3" json:"gauge,omitempty"`
	Histogram   *HistogramData   `protobuf:"bytes,4,opt,name=histogram,proto3" json:"histogram,omitempty"`
	Labels      []*LabelPair     `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty"`
	Unit        Unit             `protobuf:"varint,6,opt,name=unit,proto3,enum=metrics.Unit" json:"unit,omitempty"`
	TimestampNs int64            `protobuf:"varint,7,opt,name=TimestampNs,proto3" json:"TimestampNs,omitempty"`
}

func (x *MetricData) Reset() {
	*x = MetricData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricData) ProtoMessage() {}

func (x *MetricData) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricData.ProtoReflect.Descriptor instead.
func (*MetricData) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{3}
}

func (x *MetricData) GetCounter() *CounterData {
	if x != nil {
		return x.Counter
	}
	return nil
}

func (x *MetricData) GetCounterRate() *CounterRateData {
	if x != nil {
		return x.CounterRate
	}
	return nil
}

func (x *MetricData) GetGauge() *GaugeData {
	if x != nil {
		return x.Gauge
	}
	return nil
}

func (x *MetricData) GetHistogram() *HistogramData {
	if x != nil {
		return x.Histogram
	}
	return nil
}

func (x *MetricData) GetLabels() []*LabelPair {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *MetricData) GetUnit() Unit {
	if x != nil {
		return x.Unit
	}
	return Unit_UNSET
}

func (x *MetricData) GetTimestampNs() int64 {
	if x != nil {
		return x.TimestampNs
	}
	return 0
}

type CounterData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CounterData) Reset() {
	*x = CounterData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CounterData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterData) ProtoMessage() {}

func (x *CounterData) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CounterData.ProtoReflect.Descriptor instead.
func (*CounterData) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{4}
}

func (x *CounterData) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type CounterRateData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalValue  float64 `protobuf:"fixed64,1,opt,name=total_value,json=totalValue,proto3" json:"total_value,omitempty"`
	WindowValue float64 `protobuf:"fixed64,2,opt,name=window_value,json=windowValue,proto3" json:"window_value,omitempty"`
	AvgValue    float64 `protobuf:"fixed64,3,opt,name=avg_value,json=avgValue,proto3" json:"avg_value,omitempty"`
}

func (x *CounterRateData) Reset() {
	*x = CounterRateData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CounterRateData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterRateData) ProtoMessage() {}

func (x *CounterRateData) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CounterRateData.ProtoReflect.Descriptor instead.
func (*CounterRateData) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{5}
}

func (x *CounterRateData) GetTotalValue() float64 {
	if x != nil {
		return x.TotalValue
	}
	return 0
}

func (x *CounterRateData) GetWindowValue() float64 {
	if x != nil {
		return x.WindowValue
	}
	return 0
}

func (x *CounterRateData) GetAvgValue() float64 {
	if x != nil {
		return x.AvgValue
	}
	return 0
}

type GaugeData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GaugeData) Reset() {
	*x = GaugeData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GaugeData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GaugeData) ProtoMessage() {}

func (x *GaugeData) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GaugeData.ProtoReflect.Descriptor instead.
func (*GaugeData) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{6}
}

func (x *GaugeData) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type HistogramData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SampleCount uint64        `protobuf:"varint,1,opt,name=sample_count,json=sampleCount,proto3" json:"sample_count,omitempty"`
	SampleSum   float64       `protobuf:"fixed64,2,opt,name=sample_sum,json=sampleSum,proto3" json:"sample_sum,omitempty"`
	Buckets     []*Bucket     `protobuf:"bytes,3,rep,name=buckets,proto3" json:"buckets,omitempty"`
	Pts         []*Percentile `protobuf:"bytes,4,rep,name=pts,proto3" json:"pts,omitempty"`
}

func (x *HistogramData) Reset() {
	*x = HistogramData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistogramData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistogramData) ProtoMessage() {}

func (x *HistogramData) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistogramData.ProtoReflect.Descriptor instead.
func (*HistogramData) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{7}
}

func (x *HistogramData) GetSampleCount() uint64 {
	if x != nil {
		return x.SampleCount
	}
	return 0
}

func (x *HistogramData) GetSampleSum() float64 {
	if x != nil {
		return x.SampleSum
	}
	return 0
}

func (x *HistogramData) GetBuckets() []*Bucket {
	if x != nil {
		return x.Buckets
	}
	return nil
}

func (x *HistogramData) GetPts() []*Percentile {
	if x != nil {
		return x.Pts
	}
	return nil
}

type Bucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CumulativeCount uint64  `protobuf:"varint,1,opt,name=cumulative_count,json=cumulativeCount,proto3" json:"cumulative_count,omitempty"`
	UpperBound      float64 `protobuf:"fixed64,2,opt,name=upper_bound,json=upperBound,proto3" json:"upper_bound,omitempty"`
}

func (x *Bucket) Reset() {
	*x = Bucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bucket) ProtoMessage() {}

func (x *Bucket) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bucket.ProtoReflect.Descriptor instead.
func (*Bucket) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{8}
}

func (x *Bucket) GetCumulativeCount() uint64 {
	if x != nil {
		return x.CumulativeCount
	}
	return 0
}

func (x *Bucket) GetUpperBound() float64 {
	if x != nil {
		return x.UpperBound
	}
	return 0
}

type Percentile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Unit  float64 `protobuf:"fixed64,2,opt,name=unit,proto3" json:"unit,omitempty"`
	Value float64 `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Percentile) Reset() {
	*x = Percentile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Percentile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Percentile) ProtoMessage() {}

func (x *Percentile) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Percentile.ProtoReflect.Descriptor instead.
func (*Percentile) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{9}
}

func (x *Percentile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Percentile) GetUnit() float64 {
	if x != nil {
		return x.Unit
	}
	return 0
}

func (x *Percentile) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_model_proto protoreflect.FileDescriptor

var file_model_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x35, 0x0a, 0x09, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x50,
	0x61, 0x69, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x81, 0x01,
	0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x65, 0x6c, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x65,
	0x6c, 0x70, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0d, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52,
	0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x50, 0x61, 0x69, 0x72, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x22, 0x8e, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x46, 0x61, 0x6d, 0x69,
	0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65, 0x6c, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x65, 0x6c, 0x70, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x22, 0xca, 0x02, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x2e, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x12, 0x3b, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x61, 0x74, 0x65, 0x12, 0x28,
	0x0a, 0x05, 0x67, 0x61, 0x75, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x47, 0x61, 0x75, 0x67, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x05, 0x67, 0x61, 0x75, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x2a,
	0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x50, 0x61,
	0x69, 0x72, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x6e,
	0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4e, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4e, 0x73, 0x22,
	0x23, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x72, 0x0a, 0x0f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52,
	0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b,
	0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61,
	0x76, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x61, 0x76, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x21, 0x0a, 0x09, 0x47, 0x61, 0x75, 0x67,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa3, 0x01, 0x0a, 0x0d,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x73, 0x75, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x75, 0x6d, 0x12,
	0x29, 0x0a, 0x07, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x07, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x03, 0x70, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65, 0x52, 0x03, 0x70, 0x74,
	0x73, 0x22, 0x54, 0x0a, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x63,
	0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x70, 0x65, 0x72, 0x5f,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x75, 0x70, 0x70,
	0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x4a, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x2a, 0x51, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x54, 0x59, 0x50, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x43,
	0x4f, 0x55, 0x4e, 0x54, 0x45, 0x52, 0x52, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05,
	0x47, 0x41, 0x55, 0x47, 0x45, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x49, 0x53, 0x54, 0x4f,
	0x47, 0x52, 0x41, 0x4d, 0x10, 0x04, 0x2a, 0x5d, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x09,
	0x0a, 0x05, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x59, 0x54,
	0x45, 0x53, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4f, 0x4e, 0x53, 0x54, 0x10, 0x02, 0x12,
	0x09, 0x0a, 0x05, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x41,
	0x4e, 0x4f, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x53, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x50,
	0x45, 0x52, 0x43, 0x45, 0x4e, 0x54, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x43, 0x4f,
	0x4e, 0x44, 0x53, 0x10, 0x06, 0x42, 0x0e, 0x48, 0x01, 0x5a, 0x0a, 0x2e, 0x2e, 0x2f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_proto_rawDescOnce sync.Once
	file_model_proto_rawDescData = file_model_proto_rawDesc
)

func file_model_proto_rawDescGZIP() []byte {
	file_model_proto_rawDescOnce.Do(func() {
		file_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_proto_rawDescData)
	})
	return file_model_proto_rawDescData
}

var file_model_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_model_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_model_proto_goTypes = []interface{}{
	(MetricType)(0),         // 0: metrics.MetricType
	(Unit)(0),               // 1: metrics.Unit
	(*LabelPair)(nil),       // 2: metrics.LabelPair
	(*Metadata)(nil),        // 3: metrics.Metadata
	(*MetricFamily)(nil),    // 4: metrics.MetricFamily
	(*MetricData)(nil),      // 5: metrics.MetricData
	(*CounterData)(nil),     // 6: metrics.CounterData
	(*CounterRateData)(nil), // 7: metrics.CounterRateData
	(*GaugeData)(nil),       // 8: metrics.GaugeData
	(*HistogramData)(nil),   // 9: metrics.HistogramData
	(*Bucket)(nil),          // 10: metrics.Bucket
	(*Percentile)(nil),      // 11: metrics.Percentile
}
var file_model_proto_depIdxs = []int32{
	1,  // 0: metrics.Metadata.unit:type_name -> metrics.Unit
	2,  // 1: metrics.Metadata.labels:type_name -> metrics.LabelPair
	0,  // 2: metrics.MetricFamily.type:type_name -> metrics.MetricType
	5,  // 3: metrics.MetricFamily.metrics:type_name -> metrics.MetricData
	6,  // 4: metrics.MetricData.counter:type_name -> metrics.CounterData
	7,  // 5: metrics.MetricData.counter_rate:type_name -> metrics.CounterRateData
	8,  // 6: metrics.MetricData.gauge:type_name -> metrics.GaugeData
	9,  // 7: metrics.MetricData.histogram:type_name -> metrics.HistogramData
	2,  // 8: metrics.MetricData.labels:type_name -> metrics.LabelPair
	1,  // 9: metrics.MetricData.unit:type_name -> metrics.Unit
	10, // 10: metrics.HistogramData.buckets:type_name -> metrics.Bucket
	11, // 11: metrics.HistogramData.pts:type_name -> metrics.Percentile
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_model_proto_init() }
func file_model_proto_init() {
	if File_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelPair); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricFamily); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CounterData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CounterRateData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GaugeData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistogramData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bucket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Percentile); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_proto_goTypes,
		DependencyIndexes: file_model_proto_depIdxs,
		EnumInfos:         file_model_proto_enumTypes,
		MessageInfos:      file_model_proto_msgTypes,
	}.Build()
	File_model_proto = out.File
	file_model_proto_rawDesc = nil
	file_model_proto_goTypes = nil
	file_model_proto_depIdxs = nil
}
