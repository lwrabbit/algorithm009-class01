package main

import "fmt"

func numIslands(grid [][]byte) int {
	if len(grid)==0 || len(grid[0]) == 0{
		return 0
	}
	row:=len(grid)
	col:=len(grid[0])

	uf:=NewUnionFind(row*col)
	var directions [4][2]int =[4][2]int{{0,1},{0,-1},{-1,0},{1,0}} //定义一个方向数组。用于遍历上下左右相邻的字符，用于方法二，使代码更简洁，对比方法一可以看出


	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			// if grid[i][j]=='1'{                      //方法一
			//    if i-1>=0&&grid[i-1][j]=='1'{
			//      uf.Union(i*col+j,(i-1)*col+j)

			//     //  fmt.Println(uf.count)
			//   }
			//    if i+1<row && grid[i+1][j]=='1'{
			//      uf.Union(i*col+j,(i+1)*col+j)

			//    }
			//     if j-1>=0&& grid[i][j-1]=='1'{
			//      uf.Union(i*col+j,i*col+j-1)

			//    }
			//    if j+1<col&&grid[i][j+1]=='1'{
			//      uf.Union(i*col+j,i*col+j+1)

			//    }
			// }
			//方法二
			if grid[i][j]=='0'{
				continue
			}
			for _,d := range directions{
				nr,nc:=i+d[0],j+d[1]
				if nr>=0 && nc>=0 && nr<row &&nc<col && grid[nr][nc]=='1'{
					uf.Union(i*col+j,nr*col+nc)
				}
			}
		}
	}

	return uf.count
}

func main() {
	grid := [][]byte{[]byte("11110"),[]byte("11010"),[]byte("11000"),[]byte("00000")}
	fmt.Println(numIslands(grid))
}