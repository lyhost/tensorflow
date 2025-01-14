// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/contrib/boosted_trees/proto/learner.proto

package boosted_trees

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LearnerConfig_PruningMode int32

const (
	LearnerConfig_PRUNING_MODE_UNSPECIFIED LearnerConfig_PruningMode = 0
	LearnerConfig_PRE_PRUNE                LearnerConfig_PruningMode = 1
	LearnerConfig_POST_PRUNE               LearnerConfig_PruningMode = 2
)

var LearnerConfig_PruningMode_name = map[int32]string{
	0: "PRUNING_MODE_UNSPECIFIED",
	1: "PRE_PRUNE",
	2: "POST_PRUNE",
}

var LearnerConfig_PruningMode_value = map[string]int32{
	"PRUNING_MODE_UNSPECIFIED": 0,
	"PRE_PRUNE":                1,
	"POST_PRUNE":               2,
}

func (x LearnerConfig_PruningMode) String() string {
	return proto.EnumName(LearnerConfig_PruningMode_name, int32(x))
}

func (LearnerConfig_PruningMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7a2c2093af0c402c, []int{7, 0}
}

type LearnerConfig_GrowingMode int32

const (
	LearnerConfig_GROWING_MODE_UNSPECIFIED LearnerConfig_GrowingMode = 0
	LearnerConfig_WHOLE_TREE               LearnerConfig_GrowingMode = 1
	LearnerConfig_LAYER_BY_LAYER           LearnerConfig_GrowingMode = 2
)

var LearnerConfig_GrowingMode_name = map[int32]string{
	0: "GROWING_MODE_UNSPECIFIED",
	1: "WHOLE_TREE",
	2: "LAYER_BY_LAYER",
}

var LearnerConfig_GrowingMode_value = map[string]int32{
	"GROWING_MODE_UNSPECIFIED": 0,
	"WHOLE_TREE":               1,
	"LAYER_BY_LAYER":           2,
}

func (x LearnerConfig_GrowingMode) String() string {
	return proto.EnumName(LearnerConfig_GrowingMode_name, int32(x))
}

func (LearnerConfig_GrowingMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7a2c2093af0c402c, []int{7, 1}
}

type LearnerConfig_MultiClassStrategy int32

const (
	LearnerConfig_MULTI_CLASS_STRATEGY_UNSPECIFIED LearnerConfig_MultiClassStrategy = 0
	LearnerConfig_TREE_PER_CLASS                   LearnerConfig_MultiClassStrategy = 1
	LearnerConfig_FULL_HESSIAN                     LearnerConfig_MultiClassStrategy = 2
	LearnerConfig_DIAGONAL_HESSIAN                 LearnerConfig_MultiClassStrategy = 3
)

var LearnerConfig_MultiClassStrategy_name = map[int32]string{
	0: "MULTI_CLASS_STRATEGY_UNSPECIFIED",
	1: "TREE_PER_CLASS",
	2: "FULL_HESSIAN",
	3: "DIAGONAL_HESSIAN",
}

var LearnerConfig_MultiClassStrategy_value = map[string]int32{
	"MULTI_CLASS_STRATEGY_UNSPECIFIED": 0,
	"TREE_PER_CLASS":                   1,
	"FULL_HESSIAN":                     2,
	"DIAGONAL_HESSIAN":                 3,
}

func (x LearnerConfig_MultiClassStrategy) String() string {
	return proto.EnumName(LearnerConfig_MultiClassStrategy_name, int32(x))
}

func (LearnerConfig_MultiClassStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7a2c2093af0c402c, []int{7, 2}
}

type LearnerConfig_WeakLearnerType int32

const (
	LearnerConfig_NORMAL_DECISION_TREE    LearnerConfig_WeakLearnerType = 0
	LearnerConfig_OBLIVIOUS_DECISION_TREE LearnerConfig_WeakLearnerType = 1
)

var LearnerConfig_WeakLearnerType_name = map[int32]string{
	0: "NORMAL_DECISION_TREE",
	1: "OBLIVIOUS_DECISION_TREE",
}

var LearnerConfig_WeakLearnerType_value = map[string]int32{
	"NORMAL_DECISION_TREE":    0,
	"OBLIVIOUS_DECISION_TREE": 1,
}

func (x LearnerConfig_WeakLearnerType) String() string {
	return proto.EnumName(LearnerConfig_WeakLearnerType_name, int32(x))
}

func (LearnerConfig_WeakLearnerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7a2c2093af0c402c, []int{7, 3}
}

// Tree regularization config.
type TreeRegularizationConfig struct {
	// Classic L1/L2.
	L1 float32 `protobuf:"fixed32,1,opt,name=l1,proto3" json:"l1,omitempty"`
	L2 float32 `protobuf:"fixed32,2,opt,name=l2,proto3" json:"l2,omitempty"`
	// Tree complexity penalizes overall model complexity effectively
	// limiting how deep the tree can grow in regions with small gain.
	TreeComplexity       float32  `protobuf:"fixed32,3,opt,name=tree_complexity,json=treeComplexity,proto3" json:"tree_complexity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TreeRegularizationConfig) Reset()         { *m = TreeRegularizationConfig{} }
func (m *TreeRegularizationConfig) String() string { return proto.CompactTextString(m) }
func (*TreeRegularizationConfig) ProtoMessage()    {}
func (*TreeRegularizationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a2c2093af0c402c, []int{0}
}

func (m *TreeRegularizationConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TreeRegularizationConfig.Unmarshal(m, b)
}
func (m *TreeRegularizationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TreeRegularizationConfig.Marshal(b, m, deterministic)
}
func (m *TreeRegularizationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TreeRegularizationConfig.Merge(m, src)
}
func (m *TreeRegularizationConfig) XXX_Size() int {
	return xxx_messageInfo_TreeRegularizationConfig.Size(m)
}
func (m *TreeRegularizationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TreeRegularizationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TreeRegularizationConfig proto.InternalMessageInfo

func (m *TreeRegularizationConfig) GetL1() float32 {
	if m != nil {
		return m.L1
	}
	return 0
}

func (m *TreeRegularizationConfig) GetL2() float32 {
	if m != nil {
		return m.L2
	}
	return 0
}

func (m *TreeRegularizationConfig) GetTreeComplexity() float32 {
	if m != nil {
		return m.TreeComplexity
	}
	return 0
}

// Tree constraints config.
type TreeConstraintsConfig struct {
	// Maximum depth of the trees. The default value is 6 if not specified.
	MaxTreeDepth uint32 `protobuf:"varint,1,opt,name=max_tree_depth,json=maxTreeDepth,proto3" json:"max_tree_depth,omitempty"`
	// Min hessian weight per node.
	MinNodeWeight float32 `protobuf:"fixed32,2,opt,name=min_node_weight,json=minNodeWeight,proto3" json:"min_node_weight,omitempty"`
	// Maximum number of unique features used in the tree. Zero means there is no
	// limit.
	MaxNumberOfUniqueFeatureColumns int64    `protobuf:"varint,3,opt,name=max_number_of_unique_feature_columns,json=maxNumberOfUniqueFeatureColumns,proto3" json:"max_number_of_unique_feature_columns,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *TreeConstraintsConfig) Reset()         { *m = TreeConstraintsConfig{} }
func (m *TreeConstraintsConfig) String() string { return proto.CompactTextString(m) }
func (*TreeConstraintsConfig) ProtoMessage()    {}
func (*TreeConstraintsConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a2c2093af0c402c, []int{1}
}

func (m *TreeConstraintsConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TreeConstraintsConfig.Unmarshal(m, b)
}
func (m *TreeConstraintsConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TreeConstraintsConfig.Marshal(b, m, deterministic)
}
func (m *TreeConstraintsConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TreeConstraintsConfig.Merge(m, src)
}
func (m *TreeConstraintsConfig) XXX_Size() int {
	return xxx_messageInfo_TreeConstraintsConfig.Size(m)
}
func (m *TreeConstraintsConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TreeConstraintsConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TreeConstraintsConfig proto.InternalMessageInfo

func (m *TreeConstraintsConfig) GetMaxTreeDepth() uint32 {
	if m != nil {
		return m.MaxTreeDepth
	}
	return 0
}

func (m *TreeConstraintsConfig) GetMinNodeWeight() float32 {
	if m != nil {
		return m.MinNodeWeight
	}
	return 0
}

func (m *TreeConstraintsConfig) GetMaxNumberOfUniqueFeatureColumns() int64 {
	if m != nil {
		return m.MaxNumberOfUniqueFeatureColumns
	}
	return 0
}

// LearningRateConfig describes all supported learning rate tuners.
type LearningRateConfig struct {
	// Types that are valid to be assigned to Tuner:
	//	*LearningRateConfig_Fixed
	//	*LearningRateConfig_Dropout
	//	*LearningRateConfig_LineSearch
	Tuner                isLearningRateConfig_Tuner `protobuf_oneof:"tuner"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *LearningRateConfig) Reset()         { *m = LearningRateConfig{} }
func (m *LearningRateConfig) String() string { return proto.CompactTextString(m) }
func (*LearningRateConfig) ProtoMessage()    {}
func (*LearningRateConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a2c2093af0c402c, []int{2}
}

func (m *LearningRateConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LearningRateConfig.Unmarshal(m, b)
}
func (m *LearningRateConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LearningRateConfig.Marshal(b, m, deterministic)
}
func (m *LearningRateConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LearningRateConfig.Merge(m, src)
}
func (m *LearningRateConfig) XXX_Size() int {
	return xxx_messageInfo_LearningRateConfig.Size(m)
}
func (m *LearningRateConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LearningRateConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LearningRateConfig proto.InternalMessageInfo

type isLearningRateConfig_Tuner interface {
	isLearningRateConfig_Tuner()
}

type LearningRateConfig_Fixed struct {
	Fixed *LearningRateFixedConfig `protobuf:"bytes,1,opt,name=fixed,proto3,oneof"`
}

type LearningRateConfig_Dropout struct {
	Dropout *LearningRateDropoutDrivenConfig `protobuf:"bytes,2,opt,name=dropout,proto3,oneof"`
}

type LearningRateConfig_LineSearch struct {
	LineSearch *LearningRateLineSearchConfig `protobuf:"bytes,3,opt,name=line_search,json=lineSearch,proto3,oneof"`
}

func (*LearningRateConfig_Fixed) isLearningRateConfig_Tuner() {}

func (*LearningRateConfig_Dropout) isLearningRateConfig_Tuner() {}

func (*LearningRateConfig_LineSearch) isLearningRateConfig_Tuner() {}

func (m *LearningRateConfig) GetTuner() isLearningRateConfig_Tuner {
	if m != nil {
		return m.Tuner
	}
	return nil
}

func (m *LearningRateConfig) GetFixed() *LearningRateFixedConfig {
	if x, ok := m.GetTuner().(*LearningRateConfig_Fixed); ok {
		return x.Fixed
	}
	return nil
}

func (m *LearningRateConfig) GetDropout() *LearningRateDropoutDrivenConfig {
	if x, ok := m.GetTuner().(*LearningRateConfig_Dropout); ok {
		return x.Dropout
	}
	return nil
}

func (m *LearningRateConfig) GetLineSearch() *LearningRateLineSearchConfig {
	if x, ok := m.GetTuner().(*LearningRateConfig_LineSearch); ok {
		return x.LineSearch
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LearningRateConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LearningRateConfig_Fixed)(nil),
		(*LearningRateConfig_Dropout)(nil),
		(*LearningRateConfig_LineSearch)(nil),
	}
}

// Config for a fixed learning rate.
type LearningRateFixedConfig struct {
	LearningRate         float32  `protobuf:"fixed32,1,opt,name=learning_rate,json=learningRate,proto3" json:"learning_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LearningRateFixedConfig) Reset()         { *m = LearningRateFixedConfig{} }
func (m *LearningRateFixedConfig) String() string { return proto.CompactTextString(m) }
func (*LearningRateFixedConfig) ProtoMessage()    {}
func (*LearningRateFixedConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a2c2093af0c402c, []int{3}
}

func (m *LearningRateFixedConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LearningRateFixedConfig.Unmarshal(m, b)
}
func (m *LearningRateFixedConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LearningRateFixedConfig.Marshal(b, m, deterministic)
}
func (m *LearningRateFixedConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LearningRateFixedConfig.Merge(m, src)
}
func (m *LearningRateFixedConfig) XXX_Size() int {
	return xxx_messageInfo_LearningRateFixedConfig.Size(m)
}
func (m *LearningRateFixedConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LearningRateFixedConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LearningRateFixedConfig proto.InternalMessageInfo

func (m *LearningRateFixedConfig) GetLearningRate() float32 {
	if m != nil {
		return m.LearningRate
	}
	return 0
}

// Config for a tuned learning rate.
type LearningRateLineSearchConfig struct {
	// Max learning rate. Must be strictly positive.
	MaxLearningRate float32 `protobuf:"fixed32,1,opt,name=max_learning_rate,json=maxLearningRate,proto3" json:"max_learning_rate,omitempty"`
	// Number of learning rate values to consider between [0, max_learning_rate).
	NumSteps             int32    `protobuf:"varint,2,opt,name=num_steps,json=numSteps,proto3" json:"num_steps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LearningRateLineSearchConfig) Reset()         { *m = LearningRateLineSearchConfig{} }
func (m *LearningRateLineSearchConfig) String() string { return proto.CompactTextString(m) }
func (*LearningRateLineSearchConfig) ProtoMessage()    {}
func (*LearningRateLineSearchConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a2c2093af0c402c, []int{4}
}

func (m *LearningRateLineSearchConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LearningRateLineSearchConfig.Unmarshal(m, b)
}
func (m *LearningRateLineSearchConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LearningRateLineSearchConfig.Marshal(b, m, deterministic)
}
func (m *LearningRateLineSearchConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LearningRateLineSearchConfig.Merge(m, src)
}
func (m *LearningRateLineSearchConfig) XXX_Size() int {
	return xxx_messageInfo_LearningRateLineSearchConfig.Size(m)
}
func (m *LearningRateLineSearchConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LearningRateLineSearchConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LearningRateLineSearchConfig proto.InternalMessageInfo

func (m *LearningRateLineSearchConfig) GetMaxLearningRate() float32 {
	if m != nil {
		return m.MaxLearningRate
	}
	return 0
}

func (m *LearningRateLineSearchConfig) GetNumSteps() int32 {
	if m != nil {
		return m.NumSteps
	}
	return 0
}

// When we have a sequence of trees 1, 2, 3 ... n, these essentially represent
// weights updates in functional space, and thus we can use averaging of weight
// updates to achieve better performance. For example, we can say that our final
// ensemble will be an average of ensembles of tree 1, and ensemble of tree 1
// and tree 2 etc .. ensemble of all trees.
// Note that this averaging will apply ONLY DURING PREDICTION. The training
// stays the same.
type AveragingConfig struct {
	// Types that are valid to be assigned to Config:
	//	*AveragingConfig_AverageLastNTrees
	//	*AveragingConfig_AverageLastPercentTrees
	Config               isAveragingConfig_Config `protobuf_oneof:"config"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AveragingConfig) Reset()         { *m = AveragingConfig{} }
func (m *AveragingConfig) String() string { return proto.CompactTextString(m) }
func (*AveragingConfig) ProtoMessage()    {}
func (*AveragingConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a2c2093af0c402c, []int{5}
}

func (m *AveragingConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AveragingConfig.Unmarshal(m, b)
}
func (m *AveragingConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AveragingConfig.Marshal(b, m, deterministic)
}
func (m *AveragingConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AveragingConfig.Merge(m, src)
}
func (m *AveragingConfig) XXX_Size() int {
	return xxx_messageInfo_AveragingConfig.Size(m)
}
func (m *AveragingConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AveragingConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AveragingConfig proto.InternalMessageInfo

type isAveragingConfig_Config interface {
	isAveragingConfig_Config()
}

type AveragingConfig_AverageLastNTrees struct {
	AverageLastNTrees float32 `protobuf:"fixed32,1,opt,name=average_last_n_trees,json=averageLastNTrees,proto3,oneof"`
}

type AveragingConfig_AverageLastPercentTrees struct {
	AverageLastPercentTrees float32 `protobuf:"fixed32,2,opt,name=average_last_percent_trees,json=averageLastPercentTrees,proto3,oneof"`
}

func (*AveragingConfig_AverageLastNTrees) isAveragingConfig_Config() {}

func (*AveragingConfig_AverageLastPercentTrees) isAveragingConfig_Config() {}

func (m *AveragingConfig) GetConfig() isAveragingConfig_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *AveragingConfig) GetAverageLastNTrees() float32 {
	if x, ok := m.GetConfig().(*AveragingConfig_AverageLastNTrees); ok {
		return x.AverageLastNTrees
	}
	return 0
}

func (m *AveragingConfig) GetAverageLastPercentTrees() float32 {
	if x, ok := m.GetConfig().(*AveragingConfig_AverageLastPercentTrees); ok {
		return x.AverageLastPercentTrees
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AveragingConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AveragingConfig_AverageLastNTrees)(nil),
		(*AveragingConfig_AverageLastPercentTrees)(nil),
	}
}

type LearningRateDropoutDrivenConfig struct {
	// Probability of dropping each tree in an existing so far ensemble.
	DropoutProbability float32 `protobuf:"fixed32,1,opt,name=dropout_probability,json=dropoutProbability,proto3" json:"dropout_probability,omitempty"`
	// When trees are built after dropout happen, they don't "advance" to the
	// optimal solution, they just rearrange the path. However you can still
	// choose to skip dropout periodically, to allow a new tree that "advances"
	// to be added.
	// For example, if running for 200 steps with probability of dropout 1/100,
	// you would expect the dropout to start happening for sure for all iterations
	// after 100. However you can add probability_of_skipping_dropout of 0.1, this
	// way iterations 100-200 will include approx 90 iterations of dropout and 10
	// iterations of normal steps.Set it to 0 if you want just keep building
	// the refinement trees after dropout kicks in.
	ProbabilityOfSkippingDropout float32 `protobuf:"fixed32,2,opt,name=probability_of_skipping_dropout,json=probabilityOfSkippingDropout,proto3" json:"probability_of_skipping_dropout,omitempty"`
	// Between 0 and 1.
	LearningRate         float32  `protobuf:"fixed32,3,opt,name=learning_rate,json=learningRate,proto3" json:"learning_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LearningRateDropoutDrivenConfig) Reset()         { *m = LearningRateDropoutDrivenConfig{} }
func (m *LearningRateDropoutDrivenConfig) String() string { return proto.CompactTextString(m) }
func (*LearningRateDropoutDrivenConfig) ProtoMessage()    {}
func (*LearningRateDropoutDrivenConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a2c2093af0c402c, []int{6}
}

func (m *LearningRateDropoutDrivenConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LearningRateDropoutDrivenConfig.Unmarshal(m, b)
}
func (m *LearningRateDropoutDrivenConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LearningRateDropoutDrivenConfig.Marshal(b, m, deterministic)
}
func (m *LearningRateDropoutDrivenConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LearningRateDropoutDrivenConfig.Merge(m, src)
}
func (m *LearningRateDropoutDrivenConfig) XXX_Size() int {
	return xxx_messageInfo_LearningRateDropoutDrivenConfig.Size(m)
}
func (m *LearningRateDropoutDrivenConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LearningRateDropoutDrivenConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LearningRateDropoutDrivenConfig proto.InternalMessageInfo

func (m *LearningRateDropoutDrivenConfig) GetDropoutProbability() float32 {
	if m != nil {
		return m.DropoutProbability
	}
	return 0
}

func (m *LearningRateDropoutDrivenConfig) GetProbabilityOfSkippingDropout() float32 {
	if m != nil {
		return m.ProbabilityOfSkippingDropout
	}
	return 0
}

func (m *LearningRateDropoutDrivenConfig) GetLearningRate() float32 {
	if m != nil {
		return m.LearningRate
	}
	return 0
}

type LearnerConfig struct {
	// Number of classes.
	NumClasses uint32 `protobuf:"varint,1,opt,name=num_classes,json=numClasses,proto3" json:"num_classes,omitempty"`
	// Fraction of features to consider in each tree sampled randomly
	// from all available features.
	//
	// Types that are valid to be assigned to FeatureFraction:
	//	*LearnerConfig_FeatureFractionPerTree
	//	*LearnerConfig_FeatureFractionPerLevel
	FeatureFraction isLearnerConfig_FeatureFraction `protobuf_oneof:"feature_fraction"`
	// Regularization.
	Regularization *TreeRegularizationConfig `protobuf:"bytes,4,opt,name=regularization,proto3" json:"regularization,omitempty"`
	// Constraints.
	Constraints *TreeConstraintsConfig `protobuf:"bytes,5,opt,name=constraints,proto3" json:"constraints,omitempty"`
	// Pruning. POST_PRUNE is the default pruning mode.
	PruningMode LearnerConfig_PruningMode `protobuf:"varint,8,opt,name=pruning_mode,json=pruningMode,proto3,enum=tensorflow.boosted_trees.learner.LearnerConfig_PruningMode" json:"pruning_mode,omitempty"`
	// Growing Mode. LAYER_BY_LAYER is the default growing mode.
	GrowingMode LearnerConfig_GrowingMode `protobuf:"varint,9,opt,name=growing_mode,json=growingMode,proto3,enum=tensorflow.boosted_trees.learner.LearnerConfig_GrowingMode" json:"growing_mode,omitempty"`
	// Learning rate. By default we use fixed learning rate of 0.1.
	LearningRateTuner *LearningRateConfig `protobuf:"bytes,6,opt,name=learning_rate_tuner,json=learningRateTuner,proto3" json:"learning_rate_tuner,omitempty"`
	// Multi-class strategy. By default we use TREE_PER_CLASS for binary
	// classification and linear regression. For other cases, we use
	// DIAGONAL_HESSIAN as the default.
	MultiClassStrategy LearnerConfig_MultiClassStrategy `protobuf:"varint,10,opt,name=multi_class_strategy,json=multiClassStrategy,proto3,enum=tensorflow.boosted_trees.learner.LearnerConfig_MultiClassStrategy" json:"multi_class_strategy,omitempty"`
	// If you want to average the ensembles (for regularization), provide the
	// config below.
	AveragingConfig *AveragingConfig `protobuf:"bytes,11,opt,name=averaging_config,json=averagingConfig,proto3" json:"averaging_config,omitempty"`
	// By default we use NORMAL_DECISION_TREE as weak learner.
	WeakLearnerType      LearnerConfig_WeakLearnerType `protobuf:"varint,12,opt,name=weak_learner_type,json=weakLearnerType,proto3,enum=tensorflow.boosted_trees.learner.LearnerConfig_WeakLearnerType" json:"weak_learner_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *LearnerConfig) Reset()         { *m = LearnerConfig{} }
func (m *LearnerConfig) String() string { return proto.CompactTextString(m) }
func (*LearnerConfig) ProtoMessage()    {}
func (*LearnerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a2c2093af0c402c, []int{7}
}

func (m *LearnerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LearnerConfig.Unmarshal(m, b)
}
func (m *LearnerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LearnerConfig.Marshal(b, m, deterministic)
}
func (m *LearnerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LearnerConfig.Merge(m, src)
}
func (m *LearnerConfig) XXX_Size() int {
	return xxx_messageInfo_LearnerConfig.Size(m)
}
func (m *LearnerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LearnerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LearnerConfig proto.InternalMessageInfo

func (m *LearnerConfig) GetNumClasses() uint32 {
	if m != nil {
		return m.NumClasses
	}
	return 0
}

type isLearnerConfig_FeatureFraction interface {
	isLearnerConfig_FeatureFraction()
}

type LearnerConfig_FeatureFractionPerTree struct {
	FeatureFractionPerTree float32 `protobuf:"fixed32,2,opt,name=feature_fraction_per_tree,json=featureFractionPerTree,proto3,oneof"`
}

type LearnerConfig_FeatureFractionPerLevel struct {
	FeatureFractionPerLevel float32 `protobuf:"fixed32,3,opt,name=feature_fraction_per_level,json=featureFractionPerLevel,proto3,oneof"`
}

func (*LearnerConfig_FeatureFractionPerTree) isLearnerConfig_FeatureFraction() {}

func (*LearnerConfig_FeatureFractionPerLevel) isLearnerConfig_FeatureFraction() {}

func (m *LearnerConfig) GetFeatureFraction() isLearnerConfig_FeatureFraction {
	if m != nil {
		return m.FeatureFraction
	}
	return nil
}

func (m *LearnerConfig) GetFeatureFractionPerTree() float32 {
	if x, ok := m.GetFeatureFraction().(*LearnerConfig_FeatureFractionPerTree); ok {
		return x.FeatureFractionPerTree
	}
	return 0
}

func (m *LearnerConfig) GetFeatureFractionPerLevel() float32 {
	if x, ok := m.GetFeatureFraction().(*LearnerConfig_FeatureFractionPerLevel); ok {
		return x.FeatureFractionPerLevel
	}
	return 0
}

func (m *LearnerConfig) GetRegularization() *TreeRegularizationConfig {
	if m != nil {
		return m.Regularization
	}
	return nil
}

func (m *LearnerConfig) GetConstraints() *TreeConstraintsConfig {
	if m != nil {
		return m.Constraints
	}
	return nil
}

func (m *LearnerConfig) GetPruningMode() LearnerConfig_PruningMode {
	if m != nil {
		return m.PruningMode
	}
	return LearnerConfig_PRUNING_MODE_UNSPECIFIED
}

func (m *LearnerConfig) GetGrowingMode() LearnerConfig_GrowingMode {
	if m != nil {
		return m.GrowingMode
	}
	return LearnerConfig_GROWING_MODE_UNSPECIFIED
}

func (m *LearnerConfig) GetLearningRateTuner() *LearningRateConfig {
	if m != nil {
		return m.LearningRateTuner
	}
	return nil
}

func (m *LearnerConfig) GetMultiClassStrategy() LearnerConfig_MultiClassStrategy {
	if m != nil {
		return m.MultiClassStrategy
	}
	return LearnerConfig_MULTI_CLASS_STRATEGY_UNSPECIFIED
}

func (m *LearnerConfig) GetAveragingConfig() *AveragingConfig {
	if m != nil {
		return m.AveragingConfig
	}
	return nil
}

func (m *LearnerConfig) GetWeakLearnerType() LearnerConfig_WeakLearnerType {
	if m != nil {
		return m.WeakLearnerType
	}
	return LearnerConfig_NORMAL_DECISION_TREE
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LearnerConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LearnerConfig_FeatureFractionPerTree)(nil),
		(*LearnerConfig_FeatureFractionPerLevel)(nil),
	}
}

func init() {
	proto.RegisterEnum("tensorflow.boosted_trees.learner.LearnerConfig_PruningMode", LearnerConfig_PruningMode_name, LearnerConfig_PruningMode_value)
	proto.RegisterEnum("tensorflow.boosted_trees.learner.LearnerConfig_GrowingMode", LearnerConfig_GrowingMode_name, LearnerConfig_GrowingMode_value)
	proto.RegisterEnum("tensorflow.boosted_trees.learner.LearnerConfig_MultiClassStrategy", LearnerConfig_MultiClassStrategy_name, LearnerConfig_MultiClassStrategy_value)
	proto.RegisterEnum("tensorflow.boosted_trees.learner.LearnerConfig_WeakLearnerType", LearnerConfig_WeakLearnerType_name, LearnerConfig_WeakLearnerType_value)
	proto.RegisterType((*TreeRegularizationConfig)(nil), "tensorflow.boosted_trees.learner.TreeRegularizationConfig")
	proto.RegisterType((*TreeConstraintsConfig)(nil), "tensorflow.boosted_trees.learner.TreeConstraintsConfig")
	proto.RegisterType((*LearningRateConfig)(nil), "tensorflow.boosted_trees.learner.LearningRateConfig")
	proto.RegisterType((*LearningRateFixedConfig)(nil), "tensorflow.boosted_trees.learner.LearningRateFixedConfig")
	proto.RegisterType((*LearningRateLineSearchConfig)(nil), "tensorflow.boosted_trees.learner.LearningRateLineSearchConfig")
	proto.RegisterType((*AveragingConfig)(nil), "tensorflow.boosted_trees.learner.AveragingConfig")
	proto.RegisterType((*LearningRateDropoutDrivenConfig)(nil), "tensorflow.boosted_trees.learner.LearningRateDropoutDrivenConfig")
	proto.RegisterType((*LearnerConfig)(nil), "tensorflow.boosted_trees.learner.LearnerConfig")
}

func init() {
	proto.RegisterFile("tensorflow/contrib/boosted_trees/proto/learner.proto", fileDescriptor_7a2c2093af0c402c)
}

var fileDescriptor_7a2c2093af0c402c = []byte{
	// 1103 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdb, 0x6e, 0x23, 0x45,
	0x13, 0x8e, 0x9d, 0x3f, 0x7b, 0x28, 0x27, 0xb6, 0xd3, 0x9b, 0x9f, 0x1d, 0xd8, 0x95, 0x12, 0x99,
	0x15, 0xac, 0xb8, 0xb0, 0x95, 0xb0, 0x12, 0x82, 0x15, 0x8b, 0x7c, 0x98, 0x24, 0x46, 0x13, 0x8f,
	0xe9, 0xb1, 0x89, 0x82, 0x80, 0x56, 0xdb, 0x6e, 0x4f, 0x46, 0x99, 0x13, 0x3d, 0x3d, 0x89, 0xc3,
	0x3b, 0xf0, 0x32, 0x88, 0x6b, 0xee, 0x78, 0x27, 0x2e, 0x51, 0xf7, 0x74, 0xec, 0x89, 0x93, 0x90,
	0x0d, 0x77, 0xd3, 0xd5, 0x5f, 0x7d, 0x55, 0xf5, 0x55, 0x75, 0x69, 0xe0, 0x8d, 0x60, 0x61, 0x12,
	0xf1, 0xa9, 0x1f, 0x5d, 0x34, 0xc6, 0x51, 0x28, 0xb8, 0x37, 0x6a, 0x8c, 0xa2, 0x28, 0x11, 0x6c,
	0x42, 0x04, 0x67, 0x2c, 0x69, 0xc4, 0x3c, 0x12, 0x51, 0xc3, 0x67, 0x94, 0x87, 0x8c, 0xd7, 0xd5,
	0x09, 0xed, 0x2c, 0xbc, 0xea, 0xd7, 0xd0, 0x75, 0x8d, 0xab, 0x8d, 0xc1, 0x18, 0x70, 0xc6, 0x30,
	0x73, 0x53, 0x9f, 0x72, 0xef, 0x57, 0x2a, 0xbc, 0x28, 0x6c, 0x47, 0xe1, 0xd4, 0x73, 0x51, 0x19,
	0x8a, 0xfe, 0xae, 0x51, 0xd8, 0x29, 0xbc, 0x2e, 0xe2, 0xa2, 0xbf, 0xab, 0xce, 0x7b, 0x46, 0x51,
	0x9f, 0xf7, 0xd0, 0xa7, 0x50, 0x91, 0x64, 0x64, 0x1c, 0x05, 0xb1, 0xcf, 0x66, 0x9e, 0xb8, 0x34,
	0x56, 0xd5, 0x65, 0x59, 0x9a, 0xdb, 0x73, 0x6b, 0xed, 0x8f, 0x02, 0xfc, 0x7f, 0xa0, 0x4c, 0x61,
	0x22, 0x38, 0xf5, 0x42, 0x91, 0xe8, 0x10, 0xaf, 0xa0, 0x1c, 0xd0, 0x99, 0xca, 0x89, 0x4c, 0x58,
	0x2c, 0x4e, 0x55, 0xb8, 0x0d, 0xbc, 0x1e, 0xd0, 0x99, 0xf4, 0xe8, 0x48, 0x1b, 0xfa, 0x04, 0x2a,
	0x81, 0x17, 0x92, 0x30, 0x9a, 0x30, 0x72, 0xc1, 0x3c, 0xf7, 0x54, 0xe8, 0x2c, 0x36, 0x02, 0x2f,
	0xec, 0x45, 0x13, 0x76, 0xac, 0x8c, 0xe8, 0x08, 0x5e, 0x49, 0xb6, 0x30, 0x0d, 0x46, 0x8c, 0x93,
	0x68, 0x4a, 0xd2, 0xd0, 0xfb, 0x25, 0x65, 0x64, 0xca, 0xa8, 0x48, 0xb9, 0x4c, 0xd4, 0x4f, 0x83,
	0x30, 0x51, 0x59, 0xae, 0xe2, 0xed, 0x80, 0xce, 0x7a, 0x0a, 0x6a, 0x4f, 0x87, 0x0a, 0xb8, 0x9f,
	0xe1, 0xda, 0x19, 0xac, 0xf6, 0x7b, 0x11, 0x90, 0x25, 0x75, 0xf2, 0x42, 0x17, 0x53, 0xc1, 0x74,
	0xce, 0xdf, 0xc1, 0xda, 0xd4, 0x9b, 0xb1, 0x89, 0x4a, 0xb5, 0xb4, 0xf7, 0x65, 0xfd, 0x3e, 0x91,
	0xeb, 0x79, 0x92, 0x7d, 0xe9, 0x9a, 0x31, 0x1d, 0xae, 0xe0, 0x8c, 0x09, 0xfd, 0x04, 0x8f, 0x27,
	0x3c, 0x8a, 0xa3, 0x34, 0x2b, 0xac, 0xb4, 0xd7, 0x7c, 0x18, 0x69, 0x27, 0x73, 0xee, 0x70, 0xef,
	0x9c, 0x85, 0x73, 0xf2, 0x2b, 0x4e, 0x44, 0xa1, 0xe4, 0x7b, 0x21, 0x23, 0x09, 0xa3, 0x7c, 0x7c,
	0xaa, 0xca, 0x2f, 0xed, 0xbd, 0x7b, 0x58, 0x08, 0xcb, 0x0b, 0x99, 0xa3, 0xfc, 0xe7, 0xfc, 0xe0,
	0xcf, 0x6d, 0xad, 0xc7, 0xb0, 0x26, 0x52, 0x39, 0x50, 0xef, 0xe0, 0xf9, 0x1d, 0xe5, 0xa2, 0x8f,
	0x61, 0xc3, 0xd7, 0x57, 0x84, 0x53, 0xc1, 0xf4, 0x68, 0xad, 0xfb, 0x39, 0x7c, 0xcd, 0x85, 0x97,
	0xff, 0x16, 0x16, 0x7d, 0x06, 0x9b, 0xb2, 0xc7, 0xb7, 0x11, 0x55, 0x02, 0x3a, 0xcb, 0xfb, 0xa2,
	0x17, 0xf0, 0x34, 0x4c, 0x03, 0x92, 0x08, 0x16, 0x27, 0x4a, 0xd8, 0x35, 0xfc, 0x24, 0x4c, 0x03,
	0x47, 0x9e, 0x6b, 0xbf, 0x15, 0xa0, 0xd2, 0x3c, 0x67, 0x9c, 0xba, 0x5e, 0xe8, 0x6a, 0xf2, 0x5d,
	0xd8, 0xa2, 0xca, 0xc4, 0x88, 0x4f, 0x13, 0x41, 0xc2, 0x4c, 0x90, 0x8c, 0xff, 0x70, 0x05, 0x6f,
	0xea, 0x5b, 0x8b, 0x26, 0xa2, 0x27, 0xe7, 0x33, 0x41, 0x5f, 0xc3, 0x47, 0xd7, 0x5c, 0x62, 0xc6,
	0xc7, 0x2c, 0x14, 0xda, 0xb1, 0xa8, 0x1d, 0x9f, 0xe7, 0x1c, 0xfb, 0x19, 0x42, 0xb9, 0xb7, 0x9e,
	0xc0, 0xa3, 0xb1, 0x8a, 0x5d, 0xfb, 0xb3, 0x00, 0xdb, 0xf7, 0xf4, 0x14, 0x35, 0xe0, 0x99, 0xee,
	0x29, 0x89, 0x79, 0x34, 0xa2, 0x23, 0xcf, 0x97, 0xaf, 0x2e, 0x2b, 0x1f, 0xe9, 0xab, 0xfe, 0xe2,
	0x06, 0x99, 0xb0, 0x9d, 0x03, 0xca, 0x27, 0x91, 0x9c, 0x79, 0x71, 0x2c, 0x85, 0xcb, 0x0f, 0x5c,
	0x11, 0xbf, 0xcc, 0xc1, 0xec, 0xa9, 0xa3, 0x41, 0x3a, 0x87, 0x9b, 0x9d, 0x5b, 0xbd, 0xa5, 0x73,
	0x7f, 0x01, 0x6c, 0x58, 0xd9, 0xe4, 0xe8, 0x74, 0xb7, 0xa1, 0x24, 0xf5, 0x1f, 0xfb, 0x34, 0x49,
	0xb4, 0x8a, 0x1b, 0x18, 0xc2, 0x34, 0x68, 0x67, 0x16, 0xf4, 0x16, 0x3e, 0xbc, 0x7a, 0x9b, 0x53,
	0x4e, 0xc7, 0x72, 0xf7, 0x48, 0x01, 0x95, 0x78, 0x73, 0xed, 0x3e, 0xd0, 0x90, 0x7d, 0x8d, 0xe8,
	0x33, 0x2e, 0xb5, 0x93, 0xca, 0xdf, 0xea, 0xec, 0xb3, 0x73, 0xe6, 0x67, 0x19, 0x4a, 0xe5, 0x6f,
	0x7a, 0x5b, 0x12, 0x80, 0x46, 0x50, 0xe6, 0xd7, 0xb6, 0x9e, 0xf1, 0x3f, 0xf5, 0x2e, 0xbe, 0xba,
	0xff, 0x5d, 0xdc, 0xb5, 0x31, 0xf1, 0x12, 0x23, 0x3a, 0x81, 0xd2, 0x78, 0xb1, 0xf3, 0x8c, 0x35,
	0x15, 0xe0, 0x8b, 0xf7, 0x0b, 0x70, 0x63, 0x59, 0xe2, 0x3c, 0x17, 0xfa, 0x19, 0xd6, 0x63, 0x9e,
	0xaa, 0x8e, 0x04, 0xd1, 0x84, 0x19, 0x4f, 0x76, 0x0a, 0xaf, 0xcb, 0x7b, 0x6f, 0xdf, 0xf3, 0x51,
	0x5f, 0xb5, 0xa8, 0xde, 0xcf, 0x38, 0x8e, 0xa2, 0x09, 0xc3, 0xa5, 0x78, 0x71, 0x90, 0xfc, 0x2e,
	0x8f, 0x2e, 0xe6, 0xfc, 0x4f, 0xff, 0x1b, 0xff, 0x41, 0xc6, 0x91, 0xf1, 0xbb, 0x8b, 0x03, 0x9a,
	0xc0, 0xb3, 0x6b, 0x23, 0x45, 0xd4, 0xfa, 0x30, 0x1e, 0x29, 0x89, 0xde, 0x3c, 0x6c, 0x37, 0x69,
	0x7d, 0x36, 0xf3, 0xe3, 0x38, 0x90, 0x74, 0x48, 0xc0, 0x56, 0x90, 0xfa, 0xc2, 0xcb, 0x66, 0x90,
	0x48, 0xf5, 0x04, 0x73, 0x2f, 0x0d, 0x50, 0xd5, 0xb4, 0x1e, 0x5a, 0xcd, 0x91, 0xe4, 0x52, 0xc3,
	0xeb, 0x68, 0x26, 0x8c, 0x82, 0x1b, 0x36, 0xf4, 0x23, 0x54, 0xe9, 0xd5, 0x66, 0x21, 0xd9, 0xf3,
	0x36, 0x4a, 0xaa, 0xb0, 0xdd, 0xfb, 0x23, 0x2e, 0xed, 0x24, 0x5c, 0xa1, 0x4b, 0x4b, 0xea, 0x0c,
	0x36, 0x2f, 0x18, 0x3d, 0x23, 0xda, 0x81, 0x88, 0xcb, 0x98, 0x19, 0xeb, 0xaa, 0xa0, 0x6f, 0x1e,
	0x5a, 0xd0, 0x31, 0xa3, 0x67, 0xda, 0x32, 0xb8, 0x8c, 0x19, 0xae, 0x5c, 0x5c, 0x37, 0xd4, 0xbe,
	0x85, 0x52, 0x6e, 0x44, 0xd0, 0x4b, 0x30, 0xfa, 0x78, 0xd8, 0xeb, 0xf6, 0x0e, 0xc8, 0x91, 0xdd,
	0x31, 0xc9, 0xb0, 0xe7, 0xf4, 0xcd, 0x76, 0x77, 0xbf, 0x6b, 0x76, 0xaa, 0x2b, 0x68, 0x03, 0x9e,
	0xf6, 0xb1, 0x49, 0x24, 0xc2, 0xac, 0x16, 0x50, 0x19, 0xa0, 0x6f, 0x3b, 0x03, 0x7d, 0x2e, 0xd6,
	0x6c, 0x28, 0xe5, 0xc6, 0x41, 0x72, 0x1d, 0x60, 0xfb, 0xf8, 0x0e, 0xae, 0x32, 0xc0, 0xf1, 0xa1,
	0x6d, 0x99, 0x64, 0x80, 0x4d, 0x49, 0x86, 0xa0, 0x6c, 0x35, 0x4f, 0x4c, 0x4c, 0x5a, 0x27, 0x44,
	0x7d, 0x54, 0x8b, 0xb5, 0x73, 0x40, 0x37, 0x3b, 0x82, 0x5e, 0xc1, 0xce, 0xd1, 0xd0, 0x1a, 0x74,
	0x49, 0xdb, 0x6a, 0x3a, 0x0e, 0x71, 0x06, 0xb8, 0x39, 0x30, 0x0f, 0x4e, 0x96, 0xf8, 0x11, 0x94,
	0x25, 0x33, 0xe9, 0x9b, 0x38, 0x03, 0x56, 0x0b, 0xa8, 0x0a, 0xeb, 0xfb, 0x43, 0xcb, 0x22, 0x87,
	0xa6, 0xe3, 0x74, 0x9b, 0xbd, 0x6a, 0x11, 0x6d, 0x41, 0xb5, 0xd3, 0x6d, 0x1e, 0xd8, 0xbd, 0xe6,
	0xc2, 0xba, 0x5a, 0x3b, 0x84, 0xca, 0x92, 0x70, 0xc8, 0x80, 0xad, 0x9e, 0x8d, 0x8f, 0x9a, 0x16,
	0xe9, 0x98, 0xed, 0xae, 0xd3, 0xb5, 0x7b, 0x59, 0xe2, 0x2b, 0xe8, 0x05, 0x3c, 0xb7, 0x5b, 0x56,
	0xf7, 0xfb, 0xae, 0x3d, 0x74, 0x96, 0x2e, 0x0b, 0x2d, 0x04, 0xd5, 0xe5, 0x1d, 0xd6, 0xea, 0xfd,
	0xd0, 0x75, 0x3d, 0x71, 0x9a, 0x8e, 0xea, 0xe3, 0x28, 0x68, 0xe4, 0xfe, 0xfb, 0x6e, 0xff, 0x74,
	0xa3, 0xc6, 0x38, 0xe2, 0xec, 0xf6, 0xbf, 0xc2, 0xbf, 0x0b, 0x85, 0xd1, 0x23, 0xf5, 0x2f, 0xf8,
	0xf9, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xea, 0x9e, 0x28, 0x43, 0x0a, 0x00, 0x00,
}
