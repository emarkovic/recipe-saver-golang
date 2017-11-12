package sessions

// Save associates the provided state data with the provided sid in the store
func Save(sid SessionID, state interface{}) error {
	return nil
}

// Get retrieves the previously saved state data for the session id,
// and populates the `state` parameter with it. This will also
// reset the data's time to live in the store.
func Get(sid SessionID, state interface{}) error {
	return nil
}

// Delete deletes all state data associated with the session id from the store.
func Delete(sid SessionID) error {
	return nil
}
