export enum MessageDirection {
    /*ClientToDll is a message direction from client to dll */
    CLIENT_2_DLL = "c2d",
    CLIENT_2_SERVER = "c2s",
    /*DllToClient is a message direction from dll to client */

    DLL_2_CLIENT = "d2c",
    /*DllToClients is a message direction from dll to all clients */
    DLL_2_CLIENTS = "d2cs",
    /*ServerToClients is a message direction from server to all clients */
    SERVER_2_CLIENTS = "s2cs"
}
