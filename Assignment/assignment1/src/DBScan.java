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
        tripdata = new ArrayList<>(); 
        clusters = new ArrayList<>(); 

    }

    //build clusters 

    //find neighborhood 

    //get distance




}
