#include <iostream>
using namespace std;

int main() {

    // This is a command 
    /*
    Hello World program in C++
    */ 

    cout << "Hello, World!" << '\n';
    cout << "Welcome to C++ programming." << endl;
    
    // Math

    
    int x = 10; //asigment
    int y = 34;
    int sum = x + y;
    cout << "The value of x is: " << x << endl;
    cout << "The value of y is: " << y << endl;
    cout << "The sum of x and y is: " << sum << endl;

    // age int (angka bulat)
    int age = 17;
    int year = 2026;
    int days = 3.5;

    cout << "The value of days is: " << days << endl;
    cout << "The value of age is: " << age << endl;
    cout << "The value of year is: " << year << endl << '\n';

    //double (angka desimal)
    double price = 19.99;
    double gpu = 2.5;
    double Temperatur = 36.5;
    string dolar = "$";
    cout << "The value of price is: " << price << dolar << endl;
    cout << "The value of gpu is: " << gpu << endl;
    cout << "The value of Temperatur is: " << Temperatur << endl;

    //singgle charakteer "Cuma bisa 1 karakter"
    char grade = 'A';
    char initial = 'B';
    cout << "The value of grade is: " << grade << endl;
    cout << "The value of initial is: " << initial << endl;

    // boolean (true atau false)
    bool student = true;
    bool power = true; 
    bool forSale= true; 
    cout << "The value of student is: " << student << endl << '\n';

    // string (Variabel menggunakan tipe data string untuk menyimpan teks "...")
    string name = "Akrom";
    string city = "Bandung City";
    string hobby = "Anime and Game";
    string country = "Indonesian";
    cout << "Hello My name is " << name << ". I live in " << city << ", " << country << "." << " My hobby is " << hobby << endl;
    
    cout << "Hello my name is " << name << ", I am " << age << " yearsold" << " I like " << hobby << ", and my grade in shcool is " << grade << endl;
    
    return 0;
}