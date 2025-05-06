```mermaid
graph TD
A[インターネット] --> B(Route 53 / DNS);
B --> C(ALB DNS名);
C --> D{Application Load Balancer<br>複数のAZに跨る};
    D --> E1[パブリックSubnet AZ-a];
    D --> E2[パブリックSubnet AZ-c];
 
    E1 --> F1[プライベートSubnet AZ-a];
    E2 --> F2[プライベートSubnet AZ-c]; 
 
    D --> G(ALBターゲットグループ);
    G --> H(ECSサービス);

    H --> I1[ECSタスク Goアプリ<br>in プライベートSubnet AZ-a];
    H --> I2[ECSタスク Goアプリ<br>in プライベートSubnet AZ-c];

    I1 --> J1(CloudWatch Logs);
    I2 --> J1;
    
    F1 --> K(NAT Gateway<br>in パブリックSubnet AZ-a);
    K --> L[インターネット<br>アウトバウンド];
    F2 --> M(NAT Gateway<br>in パブリックSubnet AZ-c); 
    M --> L;

    subgraph VPC
        subgraph アベイラビリティゾーン AZ-a
            E1;
            F1;
            K;
            I1;
        end
        subgraph アベイラビリティゾーン AZ-c %% AZ-aとは物理的に分離
            E2;
            F2;
            M;
            I2;
        end
    end

    %% スタイリング (見やすくするためのオプション)
    style A fill:#f9f,stroke:#333,stroke-width:2px;
    style B fill:#ccf,stroke:#333,stroke-width:2px;
    style C fill:#ccf,stroke:#333,stroke-width:2px;
    style D fill:#bbf,stroke:#333,stroke-width:2px;
    style E1 fill:#eef,stroke:#333,stroke-width:1px;
    style E2 fill:#eef,stroke:#333,stroke-width:1px;
    style F1 fill:#ffe,stroke:#333,stroke-width:1px;
    style F2 fill:#ffe,stroke:#333,stroke-width:1px;
    style I1 fill:#afa,stroke:#333,stroke-width:2px;
    style I2 fill:#afa,stroke:#333,stroke-width:2px;
    style K fill:#ffc,stroke:#333,stroke-width:1px;
    style M fill:#ffc,stroke:#333,stroke-width:1px;
```
