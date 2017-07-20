#include "../inc/data_base.h"

void error(const char *msg)
{
    perror(msg);
    exit(0);
}

int send_message(char   *message)
{
    int sockfd, portno, n;
    struct sockaddr_in serv_addr;
    struct hostent *server;
    int num;
	
    num = atoi(message);
    num += 714;
    message = ft_itoa(num);

    portno = atoi("9002");
    sockfd = socket(AF_INET, SOCK_STREAM, 0);
    if (sockfd < 0) 
        error("ERROR opening socket");
    server = gethostbyname("localhost");
    if (server == NULL) {
        fprintf(stderr,"ERROR, no such host\n");
        exit(0);
    }
    bzero((char *) &serv_addr, sizeof(serv_addr));
    serv_addr.sin_family = AF_INET;
    bcopy((char *)server->h_addr, 
         (char *)&serv_addr.sin_addr.s_addr,
         server->h_length);
    serv_addr.sin_port = htons(portno);
    if (connect(sockfd,(struct sockaddr *) &serv_addr,sizeof(serv_addr)) < 0) 
        error("ERROR connecting");
    n = write(sockfd, message ,strlen(message));
    if (n < 0) 
         error("ERROR writing to socket");
    close(sockfd);
    return 0;
}