package redis

type HStorage struct {
	Pool    *Pool
	SetName string
}

func (h *HStorage) Get(key string) interface{} {
	conn, err := h.Pool.Get()
	defer h.Pool.Put(conn)
	if err != nil {
		return nil
	}
	data, err := conn.Cmd("HGET", h.SetName, key).Str()
	if err != nil {
		return nil
	}
	return data
}

func (h *HStorage) Set(key string, value interface{}) {
	conn, _ := h.Pool.Get()
	defer h.Pool.Put(conn)

	conn.Cmd("HSET", h.SetName, key, value)
}
