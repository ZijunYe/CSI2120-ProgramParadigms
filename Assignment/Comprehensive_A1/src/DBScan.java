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
        for(TripRecord t: tripdata){//loopng through each trip 
            if (t.is_visited()){continue;} //do not look the repeated trip 
            //otherwise: 
            t.visit(); 
            ArrayList<TripRecord> neighbour = rangeQuery(t); //first check whether it eligble to become center points 
            if(neighbour.size() >= minPts){  //check the trip can be core point 
                c++; //create the cluster 
                Cluster newCluster = new Cluster(c,t.getPickupLocation()); 
                expandCluster(t,neighbour, newCluster); //add current core point to the cluster
                this.clusters.add(newCluster); 

             }else if(neighbour.size()< minPts){
                 t.setNoise(); 
                 continue; 
             }
        }
    }

    //method for expanding cluster and adding points to cluster 
    public void expandCluster(TripRecord t, ArrayList<TripRecord> neighbour, Cluster c){
        //if it is core point, then it can be extend further 
        //if it is non-core point, but it close to the core point, join to the cluster 
            // if non-core point, not close to core point, doing nothing 
            c.addPoint(t);
            int size = neighbour.size(); 
            for(int i = 0; i < size; i++){//loop through the neighbour array
                TripRecord n1 = neighbour.get(i); 

                if(n1.getClusterId() == -1){
                    c.addPoint(n1);
                }

                if(! (n1.is_visited())){
                    n1.visit();
                    ArrayList<TripRecord> region_of_neighbour = rangeQuery(n1); 
                    if(region_of_neighbour.size() >= minPts){ //the neighbour also is the core points 
                        neighbour.addAll(region_of_neighbour); 
                        size = size + region_of_neighbour.size(); 
                    }   
                }
            } 

    }

     //find all the points around the center points
    public ArrayList<TripRecord> rangeQuery(TripRecord t){
        ArrayList<TripRecord> r = new ArrayList<>(); 
        for(TripRecord l: this.tripdata ){ //l represent a trip 
            if(distance(t.getPickupLocation(),l.getPickupLocation()) <= Eps){
                if(!l.equals(t)){
                    r.add(l); 
                }
                
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



