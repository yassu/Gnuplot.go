Q: .git/ 以下以外の AppendCurveをAppendPElemに変換したい
find . -type d -name .git -prune -o -type f -print | xargs sed -i 's/AppendCurve/AppendPElem/g'
