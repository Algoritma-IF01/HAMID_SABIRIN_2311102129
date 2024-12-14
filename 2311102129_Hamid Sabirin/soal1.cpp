// soal 1 sequential search

#include <iostream>
using namespace std;
int main()
{
    int n = 3;
    int maksArray = 6;
    int array[maksArray] = {12, 5, 5, 3, 8, 10, 11};
    bool ketemu = false;
    cout << "data dari yang di berikan adalah {12,5,5,3,8,10,11} " << endl;
    for (int i = 0; i < maksArray; i++)
    {
        if (array[i] == n)
        {
            ketemu = true;
            cout << "data nilai " << n << " ditemukan pada indekx ke-" << i << endl;
        }
    }

    return 0;
}
