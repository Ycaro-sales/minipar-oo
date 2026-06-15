#include <sys/socket.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <stdint.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static int create_client(const char* ip, int port) {
    int sock = 0;
    if ((sock = socket(AF_INET, SOCK_STREAM, 0)) < 0) { printf("Erro na criacao do socket\n"); return -1; }
    struct sockaddr_in serv_addr;
    serv_addr.sin_family = AF_INET;
    serv_addr.sin_port = htons(port);
    if (strcmp(ip, "localhost") == 0) ip = "127.0.0.1";
    inet_pton(AF_INET, ip, &serv_addr.sin_addr);
    if (connect(sock, (struct sockaddr *)&serv_addr, sizeof(serv_addr)) < 0) { printf("Erro de conexao (Servidor offline?)\n"); return -1; }
    return sock;
}

static void channel_send(int sock, const char* op, double v1, double v2) {
    if (sock < 0) return;
    char buffer[256] = {0};
    snprintf(buffer, sizeof(buffer), "%s %f %f", op, v1, v2);
    send(sock, buffer, strlen(buffer), 0); // Envia o pacote TCP real
    memset(buffer, 0, sizeof(buffer));
    int valread = recv(sock, buffer, 255, 0); // Fica travado aguardando o servidor devolver
    if (valread > 0) printf("  [CLIENTE] Recebeu a resposta TCP: %s\n", buffer);
}

static void channel_close(int sock) {
    if (sock >= 0) close(sock);
}

static void start_server(void* callback, const char* desc, const char* ip, int port) {
    int server_fd, new_socket;
    struct sockaddr_in address;
    int opt = 1;
    int addrlen = sizeof(address);
    if ((server_fd = socket(AF_INET, SOCK_STREAM, 0)) == 0) exit(EXIT_FAILURE);
    setsockopt(server_fd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt)); // Permite reiniciar rapido
    address.sin_family = AF_INET;
    address.sin_addr.s_addr = INADDR_ANY;
    address.sin_port = htons(port);
    if (bind(server_fd, (struct sockaddr *)&address, sizeof(address)) < 0) { perror("bind"); exit(EXIT_FAILURE); }
    if (listen(server_fd, 3) < 0) { perror("listen"); exit(EXIT_FAILURE); }
    printf("Server [%s] started on %s:%d\n", desc, ip, port);
    printf("Aguardando conexoes REAIS...\n");
    while(1) {
        if ((new_socket = accept(server_fd, (struct sockaddr *)&address, (socklen_t*)&addrlen)) < 0) continue;
        printf("\n  [SERVIDOR] *** Nova Conexao TCP Aceita ***\n");
        char buffer[256];
        while(1) {
            memset(buffer, 0, sizeof(buffer));
            if (recv(new_socket, buffer, 255, 0) <= 0) break; // Sai se o cliente der close()
            printf("  [SERVIDOR] Mensagem recebida: %s\n", buffer);
            char op[16]; double v1, v2, res = 0;
            if (sscanf(buffer, "%s %lf %lf", op, &v1, &v2) == 3) {
                if (strcmp(op, "+") == 0) res = v1 + v2;
                else if (strcmp(op, "-") == 0) res = v1 - v2;
                else if (strcmp(op, "*") == 0) res = v1 * v2;
                else if (strcmp(op, "/") == 0) res = v1 / v2;
                snprintf(buffer, sizeof(buffer), "%.2f", res);
                send(new_socket, buffer, strlen(buffer), 0); // Envia resposta de volta!
                printf("  [SERVIDOR] Resposta enviada: %s\n", buffer);
            }
        }
        close(new_socket);
        printf("  [SERVIDOR] *** Cliente Desconectado ***\n");
    }
}


double calcular(char* op, double v1, double v2) {
    bool t0 = op == "+";
    if (!t0) goto L1;
    float t1 = v1 + v2;
    return t1;
    goto L0;
L1:;
    bool t2 = op == "-";
    if (!t2) goto L2;
    float t3 = v1 - v2;
    return t3;
    goto L0;
L2:;
    bool t4 = op == "*";
    if (!t4) goto L3;
    float t5 = v1 * v2;
    return t5;
    goto L0;
L3:;
    bool t6 = op == "/";
    if (!t6) goto L4;
    float t7 = v1 / v2;
    return t7;
    goto L0;
L4:;
L0:;
    return 0;
}
char* desc = "Calculadora Servidor - Digite: operacao,num1,num2";
__attribute__((constructor)) void __init_calculadora_server() {
    start_server(calcular, desc, "localhost", 5000);
}
int main() {
    printf("%s\n", "Starting Calculator Server...");
    printf("%s\n", "Server is ready and waiting for connections...");
    return 0;
}
