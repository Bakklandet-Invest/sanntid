// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o helloworld_c helloworld_c.c -lpthread
#include <pthread.h>
#include <stdio.h>
#include <time.h>


pthread_mutex_t mutex;

int i = 0;

// Note the return type: void*
void* someThreadFunction1(){
	
	for(int a = 0; a < 1000000; a++){
		pthread_mutex_lock(&mutex);
		i++;
		pthread_mutex_unlock(&mutex);
	}
	
	return NULL;
}
void* someThreadFunction2(){
	
	for(int a = 0; a < 999999; a++){
		pthread_mutex_lock(&mutex);
		i--;
		pthread_mutex_unlock(&mutex);	
	}
	return NULL;
}

int main(){
	clock_t start = clock(), diff;

	pthread_t someThread1;
	pthread_t someThread2;
	pthread_create(&someThread1, NULL, someThreadFunction1, NULL);
	pthread_create(&someThread2, NULL, someThreadFunction2, NULL);
	// Arguments to a thread would be passed here ---------^
	pthread_join(someThread1, NULL);
	pthread_join(someThread2, NULL);
	printf("%d\n",i);

	diff = clock() - start;
	int msec = diff * 1000 / CLOCKS_PER_SEC;
	printf("Time taken %d seconds %d milliseconds \n", msec/1000, msec%1000);
	return 0;
}
