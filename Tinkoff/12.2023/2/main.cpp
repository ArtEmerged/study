#include <iostream>

using namespace std;

int main()
{
    int t;
    cin >> t;
    
    for (int i = 0; i < t; i++) {
        int n, ok = 0;
        bool isCorrect = true;
        cin >> n;
        for (int j = 0; j < n; j++) {
            int temp;
            cin >> temp;
            if (temp == 1) {
                ok++;
            }
            
            if (ok > 2) {
                isCorrect = false;
            }
        }
        if (isCorrect) cout<<"Yes\n";
        else cout << "No\n";
    }
    return 0;
}
