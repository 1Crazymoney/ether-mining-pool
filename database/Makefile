NAME = database
OBJ = main.o send_data.o
INC = inc/data_base.h
CFLAGS = -Wall -Wextra -Werror
LIBFT = lib/libft.a
LIBH = lib/libft.h

all: $(NAME)

%.o: $(addprefix ./src/,%.c)
	@gcc $(CFLAGS) $(INC) -c $< $(LIBH)

$(NAME): $(OBJ)
	@gcc -g $(CFLAGS) $(OBJ) $(LIBFT) -o $(NAME)

clean:
	@rm -f $(OBJ)

fclean: clean
	@rm -f $(NAME)

re: fclean all

.PHONY: clean fclean all