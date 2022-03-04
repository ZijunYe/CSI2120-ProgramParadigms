// Project CSI2120/CSI2520
// Winter 2022
// Robert Laganiere, uottawa.ca
//the main file
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type GPScoord struct {
	lat  float64
	long float64
}

type LabelledGPScoord struct {
	GPScoord
	ID    int // point ID
	Label int // cluster ID
}
type Partition struct {
	grid []LabelledGPScoord
	x    int
	y    int
}

const N int = 4
const MinPts int = 5
const eps float64 = 0.0003
const filename string = "yellow_tripdata_2009-01-15_9h_21h_clean.csv"

func main() {

	start := time.Now()

	gps, minPt, maxPt := readCSVFile(filename)
	fmt.Printf("Number of points: %d\n", len(gps))

	minPt = GPScoord{40.7, -74.}
	maxPt = GPScoord{40.8, -73.93}

	// geographical limits
	fmt.Printf("SW:(%f , %f)\n", minPt.lat, minPt.long)
	fmt.Printf("NE:(%f , %f) \n\n", maxPt.lat, maxPt.long)

	// Parallel DBSCAN STEP 1.
	//divided big pic(get each smaller section hight & width )
	incx := (maxPt.long - minPt.long) / float64(N)
	incy := (maxPt.lat - minPt.lat) / float64(N)

	var grid [N][N][]LabelledGPScoord // a grid of GPScoord slices

	// Create the partition
	// triple loop! not very efficient, but easier to understand
	partitionSize := 0
	for j := 0; j < N; j++ {
		for i := 0; i < N; i++ {

			for _, pt := range gps {

				// is it inside the expanded grid cell
				if (pt.long >= minPt.long+float64(i)*incx-eps) && (pt.long < minPt.long+float64(i+1)*incx+eps) && (pt.lat >= minPt.lat+float64(j)*incy-eps) && (pt.lat < minPt.lat+float64(j+1)*incy+eps) {

					grid[i][j] = append(grid[i][j], pt) // add the point to this slide
					partitionSize++
				}
			}
		}
	}
	jobs := make(chan Partition)

	nthread := N * N
	var mutex sync.WaitGroup
	mutex.Add(nthread)

	for j := 0; j < nthread; j++ {
		go consumer(jobs, &mutex)
	}

	for j := 0; j < N; j++ {
		for i := 0; i < N; i++ {
			newPartition := Partition{grid[i][j], i, j}
			jobs <- newPartition

		}
	}

	close(jobs)
	//wait for consumers to terminate
	mutex.Wait()

	// Parallel DBSCAN step 3.
	// merge clusters
	// *DO NOT PROGRAM THIS STEP

	end := time.Now()
	fmt.Printf("\nExecution time: %s of %d points\n", end.Sub(start), partitionSize)
	fmt.Printf("Number of CPUs: %d", runtime.NumCPU())
}

func consumer(jobs chan Partition, done *sync.WaitGroup) {
	for {
		j, more := <-jobs

		if more {
			//calling DBScan function
			xIndex := j.x
			yIndex := j.y

			DBscan(j.grid, MinPts, eps, xIndex*10000000+yIndex*1000000)

		} else {
			done.Done()
			return
		}
	}
}

// Applies DBSCAN algorithm on LabelledGPScoord points
// LabelledGPScoord: the slice of LabelledGPScoord points
// MinPts, eps: parameters for the DBSCAN algorithm
// offset: label of first cluster (also used to identify the cluster)
// returns number of clusters found
func DBscan(coords []LabelledGPScoord, MinPts int, eps float64, offset int) (nclusters int) {

	/*// *** fake code: to be rewritten
	   time.Sleep(3)
	   nclusters=0
	   for i, pt := range coords {

		  if (i==10) {
			 nclusters++
		  }
	 	  if (i==100) {
			 nclusters++
		  }
	 	  if (i==100) {
		     break
		  }

	      pt.Label= offset+nclusters
	   }
	   // *** end of fake code.*/
	visitied := make(map[GPScoord]bool, len(coords))
	for _, value := range coords {
		if value.Label != 0 {
			continue
		}
		neighbors := findNeighbours(coords, value, eps)
		if len(neighbors) >= MinPts {
			visitied[value.GPScoord] = true
			nclusters += 1
			value.Label = nclusters
			cluster := []LabelledGPScoord{}
			cluster = expandCluster(value, cluster, neighbors, visitied)
		} else {
			visitied[value.GPScoord] = false
		}

	}

	// End of DBscan function
	// Printing the result (do not remove)
	fmt.Printf("Partition %10d : [%4d,%6d]\n", offset, nclusters, len(coords))

	return nclusters //return number of cluster we have
}

func findNeighbours(coords []LabelledGPScoord, point LabelledGPScoord, eps float64) (neighbours []LabelledGPScoord) {
	for _, p := range coords {
		if distance(p.GPScoord, point.GPScoord) <= eps {
			if p.ID != point.ID {
				neighbours = append(neighbours, p)
			}
		}
	}
	return neighbours
}

func distance(point1 GPScoord, point2 GPScoord) (distance float64) {
	x := math.Pow(point2.lat-point1.lat, 2)
	y := math.Pow(point2.long-point1.long, 2)
	distance = math.Sqrt(x + y)
	return distance
}

func expandCluster(point LabelledGPScoord, currCluster []LabelledGPScoord, neighbors []LabelledGPScoord, visited map[GPScoord]bool) []LabelledGPScoord {
	currCluster = append(currCluster, point)
	for _, p := range neighbors {
		if p.Label == 0 {
			currCluster = append(currCluster, p)
			p.Label = point.Label
		}
		isVisitied := visited[p.GPScoord]
		if !isVisitied {
			currNeighbors := findNeighbours(currCluster, p, eps)
			if len(currNeighbors) >= MinPts {
				visited[p.GPScoord] = true
				neighbors = merge(neighbors, currNeighbors)
			}

		}
	}

	return currCluster
}

func merge(list1 []LabelledGPScoord, list2 []LabelledGPScoord) []LabelledGPScoord {
	for _, p := range list2 {
		list1 = append(list1, p)
	}

	return list1
}

// reads a csv file of trip records and returns a slice of the LabelledGPScoord of the pickup locations
// and the minimum and maximum GPS coordinates
func readCSVFile(filename string) (coords []LabelledGPScoord, minPt GPScoord, maxPt GPScoord) {

	coords = make([]LabelledGPScoord, 0, 5000)

	// open csv file
	src, err := os.Open(filename)
	defer src.Close()
	if err != nil {
		panic("File not found...")
	}

	// read and skip first line
	r := csv.NewReader(src)
	record, err := r.Read()
	if err != nil {
		panic("Empty file...")
	}

	minPt.long = 1000000.
	minPt.lat = 1000000.
	maxPt.long = -1000000.
	maxPt.lat = -1000000.

	var n int = 0

	for {
		// read line
		record, err = r.Read()

		// end of file?
		if err == io.EOF {
			break
		}

		if err != nil {
			panic("Invalid file format...")
		}

		// get lattitude
		lat, err := strconv.ParseFloat(record[9], 64)
		if err != nil {
			panic("Data format error (lat)...")
		}

		// is corner point?
		if lat > maxPt.lat {
			maxPt.lat = lat
		}
		if lat < minPt.lat {
			minPt.lat = lat
		}

		// get longitude
		long, err := strconv.ParseFloat(record[8], 64)
		if err != nil {
			panic("Data format error (long)...")
		}

		// is corner point?
		if long > maxPt.long {
			maxPt.long = long
		}

		if long < minPt.long {
			minPt.long = long
		}

		// add point to the slice
		n++
		pt := GPScoord{lat, long}
		coords = append(coords, LabelledGPScoord{pt, n, 0})
	}

	return coords, minPt, maxPt
}
