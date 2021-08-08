<?php

namespace Mike Kidby;

class About extends Me
{
    public function getCurrentWorkplace(): array
    {
        return [
            'workplace' => [
                'company' => 'millMountainDigital',
                'position' => 'owner'         
            ]
        ];
    }

    public function getDailyKnowledge(): array
    {
        return [
            Php::class,
            Javascript::class,
            Laravel::class,           
            ReactNative::class,
            ScssCss::class,
            ResponsiveDesign::class,
            WordPress::class,
        ];
    }

    public function getFutureGoal(): string
    {
        return 'To contribute to open source, marketing, and digital art projects.';
    }
}
