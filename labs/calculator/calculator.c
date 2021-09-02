#include <stdio.h>
#include <stdlib.h>
#include <string.h>  
// adds/subtracts/multiplies all values that are in the *values array.
// nValues is the number of values you're reading from the array
// operator will indicate if it's an addition (1), subtraction (2) or
// multiplication (3)
long calc(int operator, int nValues, int *values) {
    int ans = values[0];
    int i;
    if (operator == 1){
        for(i = 1; i < nValues; i++){
            ans += values[i];
        }
    }else if (operator == 2){
        for(i = 1; i < nValues; i++){
            ans -= values[i];
        }
    }else if (operator == 3){
        for(i = 1; i < nValues; i++){
            ans *= values[i];
        }
    }else{
        printf("ERROR \n");
        printf("Se puede sumar, restar y multiplicar");
        return 404;
    } 

    printf("Respuesta:  %d \n", ans);

    return 0;
}

int main(int argc, char **argv) {
    int arr[argc - 1];
    int op;
    if (strcmp(argv[1], "add") == 0){
        op = 1;
    }else if (strcmp(argv[1], "sub") == 0){
        op = 2;
    }else if (strcmp(argv[1], "mult") == 0){
        op = 3;
    }else {
        op = 0;
    }

    for(int i = 2; i < argc; i++){
        arr[i - 2] = atoi(argv[i]);
    }
    
    calc(op, (argc - 2), arr);
    return -1;
}
