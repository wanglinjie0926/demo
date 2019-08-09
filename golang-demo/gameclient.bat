pushd %CD%
cd /d "bin"
.\client.exe -teamID=%1 -ip="%2" -port=%3
popd

EXIT