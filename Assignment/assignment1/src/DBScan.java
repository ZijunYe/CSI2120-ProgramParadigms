import java.util.ArrayList;

public class DBScan {
    //attributes 
    public double Eps; //what the raidus that point can reach 
    public int minPts; //how many points that point can reach
    private ArrayList<TripRecord> tripdata;
    private ArrayList<Cluster> clusters;


    //constructor 
    public DBScan(double eps, int minpts, ArrayList<TripRecord>trip){
        this.Eps = eps; 
        this.minPts = minpts; 
        this.tripdata = trip;
        this.clusters = new ArrayList<Cluster>(); 
        runDBScan(); 
    }

    public void runDBScan(){
        int c = 0; 
        for(TripRecord t: tripdata){
            if (t.is_visited()){continue;} 
            t.visit(); 
            ArrayList<TripRecord> neighbour = regionQuery(t,Eps); //first check whether it eligble to become center points 
            if(neighbour.size() >= minPts){ 
                c++; 
                Cluster newCluster = new Cluster(c,t.getPickupLocation()); 
                expandCluster(t,neighbour, newCluster,Eps,minPts); 
                this.clusters.add(newCluster); 
            }

        }
    }




    //method for expanding cluster and adding points to cluster 
    public void expandCluster(TripRecord t, ArrayList<TripRecord> neighbour, Cluster c, double eps, int minpts){
        //if it is core point, then it can be extend further 
        //if it is non-core point, but it close to the core point, join to the cluster 
            // if non-core point, not close to core point, doing nothing 

            c.addPoint(t); //t is current core point 
            int size = neighbour.size(); 
            for(int i = 0; i < size; i++){//loop through the neighbour array
                TripRecord n1 = neighbour.get(i); 
                if(! (n1.is_visited())){
                    n1.visit();
                    ArrayList<TripRecord> region_of_neighbour = regionQuery(n1,eps); 
                    if(region_of_neighbour.size() >= minpts){ //the neighbour also is the core points 
                        neighbour.addAll(region_of_neighbour); 
                        size = size + region_of_neighbour.size(); 
                
                    }   
                }
                if(n1.getClusterId() == -1){ c.addPoint(n1); }
    
            } 

    }

    //find all the points around the center points
    public ArrayList<TripRecord> regionQuery(TripRecord t, double eps){
        ArrayList<TripRecord> r = new ArrayList<>(); 
        for(TripRecord l: this.tripdata ){ //l represent a trip 
            if(distance(t.getPickupLocation(),l.getPickupLocation()) <= eps){
                r.add(l); 
            }
        }
        return r; //RETURN all the trip that starting points are close to current center point
    }


    private double distance(GPScoord point1, GPScoord point2) {
        double x = Math.pow((point2.getLatitude() - point1.getLatitude()),2); 
        double y = Math.pow((point2.getLongitude()-point1.getLongitude()),2); 
        double result = Math.sqrt(x+y); 
        return result; 
    }


    //getter 

    public ArrayList<Cluster> getClusters(){
        return this.clusters; 
    }

}



