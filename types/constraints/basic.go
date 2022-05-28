package constraints

// Basic is a parameterized type that abstracts over
// the entire class of Ordered types (the set of Go
// built-in types which respond to < <= >= > == !=
// operators) and exposes this behavior via methods
// so that they fall under the Orderer constraint.
type Basic[O Ordered] struct{ Val O }

// EQ implements Orderer[Basic[O]] for Basic[O]. Returns true if the value
// of the caller is equal to that of the parameter; otherwise returns
// false.
func (self Basic[O]) EQ(other Basic[O]) bool { return self.Val == other.Val }

// LT implements Orderer[Basic[O]] for Basic[O]. Returns true if the value
// of the caller is less than that of the parameter; otherwise returns
// false.
func (self Basic[O]) LT(other Basic[O]) bool { return self.Val < other.Val }

// GT implements Orderer[Basic[O]] for Basic[O]. Returns true if the value
// of the caller is greater than that of the parameter; otherwise returns
// false.
func (self Basic[O]) GT(other Basic[O]) bool { return self.Val > other.Val }

// LE implements Orderer[Basic[O]] for Basic[O]. Returns true if the value
// of the caller is less than or equal to that of the parameter;
// otherwise returns false.
func (self Basic[O]) LE(other Basic[O]) bool { return !self.GT(other) }

// GE implements Orderer[Basic[O]] for Basic[O]. Returns true if the value
// of the caller is greater than or equal to that of the parameter;
// otherwise returns false.
func (self Basic[O]) GE(other Basic[O]) bool { return !self.LT(other) }

// NE implements Orderer[Basic[O]] for Basic[O]. Returns true
// if the value of the caller is not equal to that of the parameter;
// otherwise returns false.
func (self Basic[O]) NE(other Basic[O]) bool { return !self.EQ(other) }
