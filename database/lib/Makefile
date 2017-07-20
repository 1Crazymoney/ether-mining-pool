# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: klongric <marvin@42.fr>                    +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2017/01/10 23:18:28 by klongric          #+#    #+#              #
#    Updated: 2017/01/25 12:20:39 by klongric         ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

NAME = libft.a
INCLUDES = -I ./includes
FLAGS = -Wall -Wextra -Werror
CC = gcc
OBJ = ft_bzero.o ft_itoa.o ft_printf.o ft_strcmp.o

all: $(NAME)

%.o: %.c
		$(CC) $(FLAGS) $(INCLUDES) -c $<

$(NAME): $(OBJ)
		ar rc $(NAME) $(OBJ)
		ranlib $(NAME)

clean:
		@/bin/rm -f $(OBJ)

fclean: clean
		@/bin/rm -f $(NAME)

re: fclean all

.PHONY: clean fclean all re 
