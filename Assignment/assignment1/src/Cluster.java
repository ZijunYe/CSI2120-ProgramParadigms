import java.util.ArrayList;

public class Cluster {
    //attributes
    private ArrayList<GPScoord> points; // what are these points(list of points)
    private  int id; //the overall cluster name 
    private GPScoord center; 

    //constructor 

    public Cluster(int cluster_id, GPScoord center_point){
        this.points = new ArrayList<>(); 
        this.id = cluster_id; 
        this.center = center_point; 
    }

    //add point method
    public void addPoint(TripRecord p){
        this.points.add(p.getPickupLocation()); 
        p.getInCluster(this.id); 
    }

    //print method 
    public void printCluster(){
        System.out.println(this.id + "| " + this.center.getLongitude() + " | " + this.center.getLatitude() + " | " + this.points.size());
    }

    //getter 
    public GPScoord getcenter(){
        return center; 
    }
    public int getSizeOfCluster(){
        return points.size(); 
    }

    public int getClusterId(){
        return id; 
    }

    public void setClusterId(GPScoord center){
        this.center = center;
    }
  
    
}
