import java.util.ArrayList;
import java.util.List;

public class Cluster {
    //attributes
    private List<GPScoord> points; // what are these points(list of points)
    private  int id; //the overall cluster name 
    private GPScoord center; 

    
    //constructor 
    public Cluster(int id){
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

   


    //get average longitude 
    /*public double getavgLat(){
        double sum = 0; 
        //looping through the point array 
        
       for(int i = 0; i < num_points; i++){
           sum += (points.get(i)).getLatitude(); 
       }
       double avgLat = sum/num_points; 

        return avgLat; 
    }


    //get average latitude 
    public double getavgLon(){
        double sum = 0; 
        for(int i = 0; i < num_points; i++){
            sum += (points.get(i)).getLongitude(); 
        }
        double avgLon = sum/num_points; 

        return avgLon; 

    }*/ 




    
}
