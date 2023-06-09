// rpc protocol

enum RID
{
    RPC_GET_ONLINE_FROM_CONN, //    <tag>go:"0" rw:"1"</tag>从负载服务器中获取在线玩家
    RPC_CONN_NOTIFY_OFFLINE, //     <tag>go:"0" rw:"1"</tag>负载服务器RPC通知玩家离线
};