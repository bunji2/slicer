## slicer

Go implementation to slice file  

### Usage

```
C:\work>slice.exe
Usage: slice.exe filepath start:stop:out...
```


### Example1

```
% slice.exe foo.dat 5:15:out.dat
5-14 saved in out.dat
```

5-14 of foo.dat is saved in out.dat.

### Example2

```
% slice.exe foo.dat 5:15:.ext1 15:25:.ext2
5-14 saved in 0.ext1
15-24 saved in 1.ext2
```

5-14 of foo.dat is saved in 0.ext1 and 15-24 of foo.dat is saved in 1.ext2.


### Sample

```
C:\work>slice MrFusion.gpjb 0:6943:.gif 6943:9727:.png 9727:26632:.jpg 26632:2791486:.bmp 2791486:2794240:.gif 2794240:2796217:.png 2796217:2813627:.jpg 2813627:5578481:.bmp 5578481:5580896:.gif 5580896:5583378:.png 5583378:5601221:.jpg 5601221:8366075:.bmp 8366075:8368830:.gif 8368830:8371932:.png 8371932::.jpg

0-6942 saved in 0.gif
6943-9726 saved in 1.png
9727-26631 saved in 2.jpg
26632-2791485 saved in 3.bmp
2791486-2794239 saved in 4.gif
2794240-2796216 saved in 5.png
2796217-2813626 saved in 6.jpg
2813627-5578480 saved in 7.bmp
5578481-5580895 saved in 8.gif
5580896-5583377 saved in 9.png
5583378-5601220 saved in 10.jpg
5601221-8366074 saved in 11.bmp
8366075-8368829 saved in 12.gif
8368830-8371931 saved in 13.png
8371932-EOF saved in 14.jpg

C:\work>ls
0.gif  10.jpg  12.gif  14.jpg  3.bmp  5.png  7.bmp  9.png
1.png  11.bmp  13.png  2.jpg   4.gif  6.jpg  8.gif
```