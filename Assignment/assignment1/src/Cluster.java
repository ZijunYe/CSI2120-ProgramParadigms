import java.util.ArrayList;
import java.util.List;

public class Cluster {
    //attributes
    private List<GPScoord> points; // what are these points(list of points)
    private  int id; //the overall cluster name 
    private GPScoord center; 

    
    //constructor 
    public Cluster(int id, GPScoord p){
        this.id = id; 
        this.points =  new ArrayList<>(); 
        this.center = null;  
    }

    //functions 

 
    public void addPoint(GPScoord p){
        points.add(p);
    }




    public GPScoord getCenter(){
        return center; 
    }

    public int getId(){
        return id; 
    }

   public String printCluster(){
       return 

   }




    
}
