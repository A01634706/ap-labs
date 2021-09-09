#include <stdio.h>
#include <stdlib.h>
#include <string.h>  
/* month_day function's prototype*/
int month_day(int year, int yearday, int *pmonth, int *pday){

    static char daytab[2][12] = {
        {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
        {31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
    };
    int i, leap, mesesTot;
    leap = year%4 == 0 && year%100 != 0 || year%400 == 0;
    mesesTot = daytab[leap][1];

    if(leap == 0){
        if (yearday < 1 || yearday > 365){
            printf("Error. \n");
            printf("Introduce un numero del 1 al 366 \n");
            return 1;
        }
    }else {
        if (yearday < 1 || yearday > 366){
            printf("Error \n");
            printf("Introduce un numero del 1 al 366 \n");
            return 1;
        }
    }
    for (i = 0; i < strlen(daytab[leap]); i++)
    {
        if (yearday <= daytab[leap][i]){
            *pday = yearday;
            *pmonth = i;
            return 1;
        }else{
            yearday -= daytab[leap][i];
        }
    }
};

int main(int argc, char **argv) {
    int dia, mes;
    static char *name[] = {
       "January", "February", "March",
       "April", "May", "June",
       "July", "August", "September",
       "October", "November", "December"
    };
    month_day(atoi(argv[1]), atoi(argv[2]), &mes, &dia);

    if(dia != 0){
        printf("%s %d, %d \n", name[mes], dia, atoi(argv[1]));
    }
    return 0;
}