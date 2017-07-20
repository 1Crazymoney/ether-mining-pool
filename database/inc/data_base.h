#ifndef DATA_BASE_H
# define DATA_BASE_H
# include <stdio.h>
# include <stdlib.h>
# include <string.h>
# include <unistd.h>
# include <sys/types.h> 
# include <sys/socket.h>
# include <netinet/in.h>
# include "../lib/libft.h"
# include <netdb.h>

int     send_message();

void    error(const char *msg);

#endif