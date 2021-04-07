package album

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/junhong91/clean-architecture-go-redis/pkg/entity"
)

//RedisRepository redis repo
type RedisRepository struct {
	pool *redis.Pool
}

//NewRedisRepository create new repository
func NewRedisRepository(p *redis.Pool) *RedisRepository {
	return &RedisRepository{
		pool: p,
	}
}

//Find a album
func (r *RedisRepository) Find(id entity.ID) (*entity.Album, error) {
	result := entity.Album{}
	conn := r.pool.Get()
	defer conn.Close()
	fmt.Println("album:" + id.String())
	values, err := redis.Values(conn.Do("HGETALL", "album:"+id.String()))
	if err != nil {
		return nil, err
	} else if len(values) == 0 {
		return nil, entity.ErrNoAlbum
	}
	err = redis.ScanStruct(values, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//Store a album
func (r *RedisRepository) Store(a *entity.Album) (entity.ID, error) {
	conn := r.pool.Get()
	defer conn.Close()
	_, err := conn.Do("HMSET", "album:"+a.ID.String(), "title", a.Title, "artist", a.Artist, "id", a.ID)
	if err != nil {
		return -1, err
	}
	return a.ID, nil
}

//Delete a album
func (r *RedisRepository) Delete(id entity.ID) error {
	conn := r.pool.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", "album:"+id.String())
	if err != nil {
		return err
	}
	return nil
}
