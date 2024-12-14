// binary search

#include <iostream>

using namespace std;
int urut = 0;
int array[7] = {2, 5, 5, 6, 8, 10, 11};

int binary()
{
    for (int i = 0; i < 7; i++)
    {
        if (array[i] == urut)
        {
            array[i] = urut;
            urut++;
        }
    }

    return 0;
}
int main()
{
    int angka;
    cout << "data dari yang di berikan adalah {2,5,5,6,8,10,11} " << endl;
    cout << "masukkan angka yang ingin dicari : ";
    cin >> angka;
    urut = angka;
    for (int i = 0; i < 7; i++)
    {
        if (array[i] == urut)
        {
            cout << "data nilai " << urut << " ditemukan pada indekx ke-" << i << endl;
        }
    }

    return 0;
}