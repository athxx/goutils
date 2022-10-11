package rdx

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

func ctx() context.Context {
	return context.Background()
}

func Pipeline() redis.Pipeliner { return rdb.Pipeline() }
func Pipelined(fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return rdb.Pipelined(ctx(), fn)
}
func TxPipelined(fn func(redis.Pipeliner) error) ([]redis.Cmder, error) { return nil, nil }
func TxPipeline() redis.Pipeliner                                       { return nil }

func Command() *redis.CommandsInfoCmd           { return nil }
func ClientGetName() *redis.StringCmd           { return nil }
func Echo(message interface{}) *redis.StringCmd { return nil }
func Ping() *redis.StatusCmd                    { return rdb.Ping(ctx()) }
func Quit() *redis.StatusCmd                    { return nil }
func Del(keys ...string) *redis.IntCmd          { return rdb.Del(ctx(), keys...) }
func Unlink(keys ...string) *redis.IntCmd       { return nil }
func Dump(key string) *redis.StringCmd          { return nil }
func Exists(keys ...string) int64 {
	return rdb.Exists(ctx(), keys...).Val()
}
func Expire(key string, expiration int) *redis.BoolCmd {
	return rdb.Expire(ctx(), key, time.Second*time.Duration(expiration))
}
func ExpireAt(key string, tm time.Time) *redis.BoolCmd { return rdb.ExpireAt(ctx(), key, tm) }
func Keys(pattern string) *redis.StringSliceCmd        { return rdb.Keys(ctx(), pattern) }
func Migrate(host, port, key string, db int, timeout int) *redis.StatusCmd {
	return rdb.Migrate(ctx(), host, port, key, db, time.Duration(timeout)*time.Second)
}

func Move(key string, db int) *redis.BoolCmd                            { return nil }
func ObjectRefCount(key string) *redis.IntCmd                           { return nil }
func ObjectEncoding(key string) *redis.StringCmd                        { return nil }
func ObjectIdleTime(key string) *redis.DurationCmd                      { return nil }
func Persist(key string) *redis.BoolCmd                                 { return nil }
func PExpire(key string, expiration int) *redis.BoolCmd                 { return nil }
func PExpireAt(key string, tm time.Time) *redis.BoolCmd                 { return nil }
func PTTL(key string) *redis.DurationCmd                                { return nil }
func RandomKey() *redis.StringCmd                                       { return nil }
func Rename(key, newkey string) *redis.StatusCmd                        { return nil }
func RenameNX(key, newkey string) *redis.BoolCmd                        { return nil }
func Restore(key string, ttl int, value string) *redis.StatusCmd        { return nil }
func RestoreReplace(key string, ttl int, value string) *redis.StatusCmd { return nil }
func Sort(key string, sort *redis.Sort) *redis.StringSliceCmd           { return nil }
func SortStore(key, store string, sort *redis.Sort) *redis.IntCmd       { return nil }
func SortInterfaces(key string, sort *redis.Sort) *redis.SliceCmd       { return nil }
func Touch(keys ...string) *redis.IntCmd                                { return nil }
func TTL(key string) *redis.DurationCmd                                 { return nil }
func Type(key string) *redis.StatusCmd                                  { return nil }
func Append(key, value string) *redis.IntCmd                            { return nil }
func Decr(key string) *redis.IntCmd                                     { return nil }
func DecrBy(key string, decrement int64) *redis.IntCmd                  { return nil }
func Get(key string) *redis.StringCmd                                   { return rdb.Get(ctx(), key) }
func GetRange(key string, start, end int64) *redis.StringCmd            { return nil }
func GetSet(key string, value interface{}) *redis.StringCmd             { return rdb.GetSet(ctx(), key, value) }
func GetEx(key string, expiration int) *redis.StringCmd {
	return rdb.GetEx(ctx(), key, time.Duration(expiration)*time.Second)
}
func GetDel(key string) *redis.StringCmd { return rdb.GetDel(ctx(), key) }
func Incr(key string) *redis.IntCmd      { return rdb.Incr(ctx(), key) }
func IncrBy(key string, value int64) *redis.IntCmd {
	return rdb.IncrBy(ctx(), key, value)
}
func IncrByFloat(key string, value float64) *redis.FloatCmd {
	return rdb.IncrByFloat(ctx(), key, value)
}
func MGet(keys ...string) *redis.SliceCmd         { return rdb.MGet(ctx(), keys...) }
func MSet(values ...interface{}) *redis.StatusCmd { return rdb.MSet(ctx(), values...) }
func MSetNX(values ...interface{}) *redis.BoolCmd { return nil }
func Set(key string, value interface{}, expiration int) *redis.StatusCmd {
	return rdb.Set(ctx(), key, value, time.Duration(expiration)*time.Second)
}
func SetArgs(key string, value interface{}, a redis.SetArgs) *redis.StatusCmd {
	return rdb.SetArgs(ctx(), key, value, a)
}

func SetEx(key string, value interface{}, expiration int) *redis.StatusCmd {
	return rdb.SetEX(ctx(), key, value, time.Duration(expiration)*time.Second)
}
func SetNX(key string, value interface{}, expiration int) *redis.BoolCmd {
	return rdb.SetNX(ctx(), key, value, time.Duration(expiration)*time.Second)
}
func SetXX(key string, value interface{}, expiration int) *redis.BoolCmd {
	return rdb.SetXX(ctx(), key, value, time.Duration(expiration)*time.Second)
}
func SetRange(key string, offset int64, value string) *redis.IntCmd { return nil }
func StrLen(key string) *redis.IntCmd                               { return nil }

func GetBit(key string, offset int64) *redis.IntCmd               { return nil }
func SetBit(key string, offset int64, value int) *redis.IntCmd    { return nil }
func BitCount(key string, bitCount *redis.BitCount) *redis.IntCmd { return nil }
func BitOpAnd(destKey string, keys ...string) *redis.IntCmd       { return nil }
func BitOpOr(destKey string, keys ...string) *redis.IntCmd        { return nil }
func BitOpXor(destKey string, keys ...string) *redis.IntCmd       { return nil }
func BitOpNot(destKey string, key string) *redis.IntCmd           { return nil }
func BitPos(key string, bit int64, pos ...int64) *redis.IntCmd    { return nil }
func BitField(key string, args ...interface{}) *redis.IntSliceCmd { return nil }

func Scan(cursor uint64, match string, count int64) *redis.ScanCmd                     { return nil }
func ScanType(cursor uint64, match string, count int64, keyType string) *redis.ScanCmd { return nil }
func SScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd        { return nil }
func HScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd        { return nil }
func ZScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd        { return nil }

func HDel(key string, fields ...string) *redis.IntCmd {
	return rdb.HDel(ctx(), key, fields...)
}
func HExists(key, field string) *redis.BoolCmd     { return rdb.HExists(ctx(), key, field) }
func HGet(key, field string) *redis.StringCmd      { return rdb.HGet(ctx(), key, field) }
func HGetAll(key string) *redis.StringStringMapCmd { return rdb.HGetAll(ctx(), key) }
func HIncrBy(key, field string, incr int64) *redis.IntCmd {
	return rdb.HIncrBy(ctx(), key, field, incr)
}
func HIncrByFloat(key, field string, incr float64) *redis.FloatCmd {
	return rdb.HIncrByFloat(ctx(), key, field, incr)
}
func HKeys(key string) *redis.StringSliceCmd { return rdb.HKeys(ctx(), key) }
func HLen(key string) *redis.IntCmd          { return rdb.HLen(ctx(), key) }
func HMGet(key string, fields ...string) *redis.SliceCmd {
	return rdb.HMGet(ctx(), key, fields...)
}
func HSet(key string, values ...interface{}) *redis.IntCmd {
	return rdb.HSet(ctx(), key, values...)
}
func HMSet(key string, values ...interface{}) *redis.BoolCmd {
	return rdb.HMSet(ctx(), key, values...)
}
func HSetNX(key, field string, value interface{}) *redis.BoolCmd              { return nil }
func HVals(key string) *redis.StringSliceCmd                                  { return nil }
func HRandField(key string, count int, withValues bool) *redis.StringSliceCmd { return nil }

func BLPop(timeout int, keys ...string) *redis.StringSliceCmd             { return nil }
func BRPop(timeout int, keys ...string) *redis.StringSliceCmd             { return nil }
func BRPopLPush(source, destination string, timeout int) *redis.StringCmd { return nil }
func LIndex(key string, index int64) *redis.StringCmd                     { return nil }
func LInsert(key, op string, pivot, value interface{}) *redis.IntCmd      { return nil }
func LInsertBefore(key string, pivot, value interface{}) *redis.IntCmd    { return nil }
func LInsertAfter(key string, pivot, value interface{}) *redis.IntCmd     { return nil }
func LLen(key string) *redis.IntCmd                                       { return nil }
func LPop(key string) *redis.StringCmd                                    { return nil }
func LPopCount(key string, count int) *redis.StringSliceCmd               { return nil }
func LPos(key string, value string, args redis.LPosArgs) *redis.IntCmd    { return nil }
func LPosCount(key string, value string, count int64, args redis.LPosArgs) *redis.IntSliceCmd {
	return nil
}
func LPush(key string, values ...interface{}) *redis.IntCmd {
	return rdb.LPush(ctx(), key, values...)
}
func LPushX(key string, values ...interface{}) *redis.IntCmd             { return nil }
func LRange(key string, start, stop int64) *redis.StringSliceCmd         { return nil }
func LRem(key string, count int64, value interface{}) *redis.IntCmd      { return nil }
func LSet(key string, index int64, value interface{}) *redis.StatusCmd   { return nil }
func LTrim(key string, start, stop int64) *redis.StatusCmd               { return nil }
func RPop(key string) *redis.StringCmd                                   { return nil }
func RPopLPush(source, destination string) *redis.StringCmd              { return nil }
func RPush(key string, values ...interface{}) *redis.IntCmd              { return nil }
func RPushX(key string, values ...interface{}) *redis.IntCmd             { return nil }
func LMove(source, destination, srcpos, destpos string) *redis.StringCmd { return nil }

func SAdd(key string, members ...interface{}) *redis.IntCmd               { return nil }
func SCard(key string) *redis.IntCmd                                      { return nil }
func SDiff(keys ...string) *redis.StringSliceCmd                          { return nil }
func SDiffStore(destination string, keys ...string) *redis.IntCmd         { return nil }
func SInter(keys ...string) *redis.StringSliceCmd                         { return nil }
func SInterStore(destination string, keys ...string) *redis.IntCmd        { return nil }
func SIsMember(key string, member interface{}) *redis.BoolCmd             { return nil }
func SMIsMember(key string, members ...interface{}) *redis.BoolSliceCmd   { return nil }
func SMembers(key string) *redis.StringSliceCmd                           { return nil }
func SMembersMap(key string) *redis.StringStructMapCmd                    { return nil }
func SMove(source, destination string, member interface{}) *redis.BoolCmd { return nil }
func SPop(key string) *redis.StringCmd                                    { return nil }
func SPopN(key string, count int64) *redis.StringSliceCmd                 { return nil }
func SRandMember(key string) *redis.StringCmd                             { return nil }
func SRandMemberN(key string, count int64) *redis.StringSliceCmd          { return nil }
func SRem(key string, members ...interface{}) *redis.IntCmd               { return nil }
func SUnion(keys ...string) *redis.StringSliceCmd                         { return nil }
func SUnionStore(destination string, keys ...string) *redis.IntCmd        { return nil }

func XAdd(a *redis.XAddArgs) *redis.StringCmd                                           { return nil }
func XDel(stream string, ids ...string) *redis.IntCmd                                   { return nil }
func XLen(stream string) *redis.IntCmd                                                  { return nil }
func XRange(stream, start, stop string) *redis.XMessageSliceCmd                         { return nil }
func XRangeN(stream, start, stop string, count int64) *redis.XMessageSliceCmd           { return nil }
func XRevRange(stream string, start, stop string) *redis.XMessageSliceCmd               { return nil }
func XRevRangeN(stream string, start, stop string, count int64) *redis.XMessageSliceCmd { return nil }
func XRead(a *redis.XReadArgs) *redis.XStreamSliceCmd                                   { return nil }
func XReadStreams(streams ...string) *redis.XStreamSliceCmd                             { return nil }
func XGroupCreate(stream, group, start string) *redis.StatusCmd                         { return nil }
func XGroupCreateMkStream(stream, group, start string) *redis.StatusCmd                 { return nil }
func XGroupSetID(stream, group, start string) *redis.StatusCmd                          { return nil }
func XGroupDestroy(stream, group string) *redis.IntCmd                                  { return nil }
func XGroupDelConsumer(stream, group, consumer string) *redis.IntCmd                    { return nil }
func XReadGroup(a *redis.XReadGroupArgs) *redis.XStreamSliceCmd                         { return nil }
func XAck(stream, group string, ids ...string) *redis.IntCmd                            { return nil }
func XPending(stream, group string) *redis.XPendingCmd                                  { return nil }
func XPendingExt(a *redis.XPendingExtArgs) *redis.XPendingExtCmd                        { return nil }
func XClaim(a *redis.XClaimArgs) *redis.XMessageSliceCmd                                { return nil }
func XClaimJustID(a *redis.XClaimArgs) *redis.StringSliceCmd                            { return nil }
func XTrim(key string, maxLen int64) *redis.IntCmd                                      { return nil }
func XTrimApprox(key string, maxLen int64) *redis.IntCmd                                { return nil }
func XInfoGroups(key string) *redis.XInfoGroupsCmd                                      { return nil }
func XInfoStream(key string) *redis.XInfoStreamCmd                                      { return nil }
func XInfoConsumers(key string, group string) *redis.XInfoConsumersCmd                  { return nil }

func BZPopMax(timeout int, keys ...string) *redis.ZWithKeyCmd                     { return nil }
func BZPopMin(timeout int, keys ...string) *redis.ZWithKeyCmd                     { return nil }
func ZAdd(key string, members ...*redis.Z) *redis.IntCmd                          { return nil }
func ZAddNX(key string, members ...*redis.Z) *redis.IntCmd                        { return nil }
func ZAddXX(key string, members ...*redis.Z) *redis.IntCmd                        { return nil }
func ZAddCh(key string, members ...*redis.Z) *redis.IntCmd                        { return nil }
func ZAddNXCh(key string, members ...*redis.Z) *redis.IntCmd                      { return nil }
func ZAddXXCh(key string, members ...*redis.Z) *redis.IntCmd                      { return nil }
func ZIncr(key string, member *redis.Z) *redis.FloatCmd                           { return nil }
func ZIncrNX(key string, member *redis.Z) *redis.FloatCmd                         { return nil }
func ZIncrXX(key string, member *redis.Z) *redis.FloatCmd                         { return nil }
func ZCard(key string) *redis.IntCmd                                              { return nil }
func ZCount(key, min, max string) *redis.IntCmd                                   { return nil }
func ZLexCount(key, min, max string) *redis.IntCmd                                { return nil }
func ZIncrBy(key string, increment float64, member string) *redis.FloatCmd        { return nil }
func ZInter(store *redis.ZStore) *redis.StringSliceCmd                            { return nil }
func ZInterWithScores(store *redis.ZStore) *redis.ZSliceCmd                       { return nil }
func ZInterStore(destination string, store *redis.ZStore) *redis.IntCmd           { return nil }
func ZMScore(key string, members ...string) *redis.FloatSliceCmd                  { return nil }
func ZPopMax(key string, count ...int64) *redis.ZSliceCmd                         { return nil }
func ZPopMin(key string, count ...int64) *redis.ZSliceCmd                         { return nil }
func ZRange(key string, start, stop int64) *redis.StringSliceCmd                  { return nil }
func ZRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd             { return nil }
func ZRangeByScore(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd         { return nil }
func ZRangeByLex(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd           { return nil }
func ZRangeByScoreWithScores(key string, opt *redis.ZRangeBy) *redis.ZSliceCmd    { return nil }
func ZRank(key, member string) *redis.IntCmd                                      { return nil }
func ZRem(key string, members ...interface{}) *redis.IntCmd                       { return nil }
func ZRemRangeByRank(key string, start, stop int64) *redis.IntCmd                 { return nil }
func ZRemRangeByScore(key, min, max string) *redis.IntCmd                         { return nil }
func ZRemRangeByLex(key, min, max string) *redis.IntCmd                           { return nil }
func ZRevRange(key string, start, stop int64) *redis.StringSliceCmd               { return nil }
func ZRevRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd          { return nil }
func ZRevRangeByScore(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd      { return nil }
func ZRevRangeByLex(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd        { return nil }
func ZRevRangeByScoreWithScores(key string, opt *redis.ZRangeBy) *redis.ZSliceCmd { return nil }
func ZRevRank(key, member string) *redis.IntCmd                                   { return nil }
func ZScore(key, member string) *redis.FloatCmd                                   { return nil }
func ZUnionStore(dest string, store *redis.ZStore) *redis.IntCmd                  { return nil }
func ZRandMember(key string, count int, withScores bool) *redis.StringSliceCmd    { return nil }
func ZDiff(keys ...string) *redis.StringSliceCmd                                  { return nil }
func ZDiffWithScores(keys ...string) *redis.ZSliceCmd                             { return nil }
func ZDiffStore(destination string, keys ...string) *redis.IntCmd                 { return nil }

func PFAdd(key string, els ...interface{}) *redis.IntCmd   { return nil }
func PFCount(keys ...string) *redis.IntCmd                 { return nil }
func PFMerge(dest string, keys ...string) *redis.StatusCmd { return nil }

func BgRewriteAOF() *redis.StatusCmd                       { return nil }
func BgSave() *redis.StatusCmd                             { return nil }
func ClientKill(ipPort string) *redis.StatusCmd            { return nil }
func ClientKillByFilter(keys ...string) *redis.IntCmd      { return nil }
func ClientList() *redis.StringCmd                         { return nil }
func ClientPause(dur int) *redis.BoolCmd                   { return nil }
func ClientID() *redis.IntCmd                              { return nil }
func ConfigGet(parameter string) *redis.SliceCmd           { return nil }
func ConfigResetStat() *redis.StatusCmd                    { return nil }
func ConfigSet(parameter, value string) *redis.StatusCmd   { return nil }
func ConfigRewrite() *redis.StatusCmd                      { return nil }
func DBSize() *redis.IntCmd                                { return nil }
func FlushAll() *redis.StatusCmd                           { return rdb.FlushAll(ctx()) }
func FlushAllAsync() *redis.StatusCmd                      { return nil }
func FlushDB() *redis.StatusCmd                            { return rdb.FlushDB(ctx()) }
func FlushDBAsync() *redis.StatusCmd                       { return nil }
func Info(section ...string) *redis.StringCmd              { return nil }
func LastSave() *redis.IntCmd                              { return nil }
func Save() *redis.StatusCmd                               { return nil }
func Shutdown() *redis.StatusCmd                           { return nil }
func ShutdownSave() *redis.StatusCmd                       { return nil }
func ShutdownNoSave() *redis.StatusCmd                     { return nil }
func SlaveOf(host, port string) *redis.StatusCmd           { return nil }
func Time() *redis.TimeCmd                                 { return nil }
func DebugObject(key string) *redis.StringCmd              { return nil }
func ReadOnly() *redis.StatusCmd                           { return nil }
func ReadWrite() *redis.StatusCmd                          { return nil }
func MemoryUsage(key string, samples ...int) *redis.IntCmd { return nil }

func Eval(script string, keys []string, args ...interface{}) *redis.Cmd  { return nil }
func EvalSha(sha1 string, keys []string, args ...interface{}) *redis.Cmd { return nil }
func ScriptExists(hashes ...string) *redis.BoolSliceCmd                  { return nil }
func ScriptFlush() *redis.StatusCmd                                      { return nil }
func ScriptKill() *redis.StatusCmd                                       { return nil }
func ScriptLoad(script string) *redis.StringCmd                          { return nil }

func Publish(channel string, message interface{}) *redis.IntCmd { return nil }
func PubSubChannels(pattern string) *redis.StringSliceCmd       { return nil }
func PubSubNumSub(channels ...string) *redis.StringIntMapCmd    { return nil }
func PubSubNumPat() *redis.IntCmd                               { return nil }

func ClusterSlots() *redis.ClusterSlotsCmd                           { return nil }
func ClusterNodes() *redis.StringCmd                                 { return nil }
func ClusterMeet(host, port string) *redis.StatusCmd                 { return nil }
func ClusterForget(nodeID string) *redis.StatusCmd                   { return nil }
func ClusterReplicate(nodeID string) *redis.StatusCmd                { return nil }
func ClusterResetSoft() *redis.StatusCmd                             { return nil }
func ClusterResetHard() *redis.StatusCmd                             { return nil }
func ClusterInfo() *redis.StringCmd                                  { return nil }
func ClusterKeySlot(key string) *redis.IntCmd                        { return nil }
func ClusterGetKeysInSlot(slot int, count int) *redis.StringSliceCmd { return nil }
func ClusterCountFailureReports(nodeID string) *redis.IntCmd         { return nil }
func ClusterCountKeysInSlot(slot int) *redis.IntCmd                  { return nil }
func ClusterDelSlots(slots ...int) *redis.StatusCmd                  { return nil }
func ClusterDelSlotsRange(min, max int) *redis.StatusCmd             { return nil }
func ClusterSaveConfig() *redis.StatusCmd                            { return nil }
func ClusterSlaves(nodeID string) *redis.StringSliceCmd              { return nil }
func ClusterFailover() *redis.StatusCmd                              { return nil }
func ClusterAddSlots(slots ...int) *redis.StatusCmd                  { return nil }
func ClusterAddSlotsRange(min, max int) *redis.StatusCmd             { return nil }

func GeoAdd(key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd { return nil }
func GeoPos(key string, members ...string) *redis.GeoPosCmd              { return nil }
func GeoRadius(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	return nil
}
func GeoRadiusStore(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.IntCmd {
	return nil
}
func GeoRadiusByMember(key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	return nil
}
func GeoRadiusByMemberStore(key, member string, query *redis.GeoRadiusQuery) *redis.IntCmd {
	return nil
}
func GeoDist(key string, member1, member2, unit string) *redis.FloatCmd { return nil }
func GeoHash(key string, members ...string) *redis.StringSliceCmd       { return nil }

func GetJson(key string, dst interface{}) error {
	data, err := Get(key).Bytes()
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, dst); err != nil {
		return err
	}
	return nil
}

func SetJson(key string, value interface{}, expiration int) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	if expiration > 0 {
		return SetEx(key, data, expiration).Err()
	}
	return Set(key, data, expiration).Err()
}
